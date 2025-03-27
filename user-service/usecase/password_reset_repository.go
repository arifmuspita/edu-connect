package usecase

import (
	"time"
	"userService/queue"
	"userService/repository"

	customErr "userService/error"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type IPasswordResetUseCase interface {
	RequestReset(email string) error
	ResetPassword(token, newPassword string) error
}

type passwordResetUseCase struct {
	userRepo       repository.IUserRepository
	resetRepo      repository.IPasswordResetRepository
	emailPublisher queue.IEmailPublisher
	logger         *logrus.Entry
}

func NewPasswordResetUseCase(
	userRepo repository.IUserRepository,
	resetRepo repository.IPasswordResetRepository,
	emailPublisher queue.IEmailPublisher,
	baseLogger *logrus.Logger,
) IPasswordResetUseCase {
	enrichedLogger := baseLogger.WithField("layer", "usecase")
	return &passwordResetUseCase{
		userRepo:       userRepo,
		resetRepo:      resetRepo,
		emailPublisher: emailPublisher,
		logger:         enrichedLogger,
	}
}

func (u *passwordResetUseCase) RequestReset(email string) error {
	u.logger.WithField("email", email)

	if email == "" {
		u.logger.Warn("Email is required")
		return customErr.ErrRegisterEmailRequired
	}

	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		u.logger.Warn("Email not found")
		return customErr.ErrLoginEmailNotFound
	}

	activeReset, err := u.resetRepo.GetActivePasswordResetByEmail(email)
	if err == nil && activeReset != nil {
		u.logger.Warn("Forgot password request denied: Active reset token exists")
		return customErr.ErrResetTokenStillValid
	}

	token := uuid.NewString()
	expiresAt := time.Now().Add(15 * time.Minute)

	err = u.resetRepo.CreateResetToken(user.Email, token, expiresAt)
	if err != nil {
		u.logger.WithError(err).Error("Failed to store reset token")
		return customErr.ErrInternalServer
	}

	err = u.emailPublisher.PublishResetPasswordToken(user.Email, token)
	if err != nil {
		u.logger.WithError(err).Error("Failed to publish reset email")
	}

	u.logger.Info("Reset password token generated and published")
	return nil
}

func (u *passwordResetUseCase) ResetPassword(token, newPassword string) error {
	u.logger.WithField("token", token)

	if token == "" {
		return customErr.ErrVerificationTokenInvalid
	}

	if !isValidPassword(newPassword) {
		return customErr.ErrRegisterInvalidPassword
	}

	data, err := u.resetRepo.ValidateResetToken(token)
	if err != nil {
		u.logger.WithError(err).Warn("Invalid or expired reset token")
		return customErr.ErrVerificationTokenInvalid
	}

	err = u.userRepo.UpdatePasswordByEmail(data.Email, newPassword)
	if err != nil {
		u.logger.WithError(err).Error("Failed to update password")
		return customErr.ErrInternalServer
	}

	err = u.resetRepo.MarkResetTokenUsed(token)
	if err != nil {
		u.logger.WithError(err).Warn("Failed to mark reset token as used")
	}

	u.logger.WithField("email", data.Email).Info("Password reset successfully")
	return nil
}
