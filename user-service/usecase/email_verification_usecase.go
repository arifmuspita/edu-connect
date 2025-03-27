package usecase

import (
	"time"
	"userService/queue"
	"userService/repository"

	customErr "userService/error"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type IVerificationUseCase interface {
	GenerateVerification(email string) error
	VerifyToken(token string) error
	ResendVerification(email string) error
}

type verificationUseCase struct {
	userRepo         repository.IUserRepository
	verificationRepo repository.IVerificationRepository
	emailPublisher   queue.IEmailPublisher
	logger           *logrus.Entry
}

func NewVerificationUseCase(
	userRepo repository.IUserRepository,
	verificationRepo repository.IVerificationRepository,
	emailPublisher queue.IEmailPublisher,
	baseLogger *logrus.Logger,
) IVerificationUseCase {
	enrichedLogger := baseLogger.WithField("layer", "usecase")
	return &verificationUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		emailPublisher:   emailPublisher,
		logger:           enrichedLogger,
	}
}

func (v *verificationUseCase) GenerateVerification(email string) error {
	v.logger.WithField("email", email)

	token := uuid.NewString()
	expiresAt := time.Now().Add(15 * time.Minute)

	err := v.verificationRepo.CreateToken(email, token, expiresAt)
	if err != nil {
		v.logger.WithError(err).Error("Failed to create email verification token")
		return customErr.ErrInternalServer
	}

	err = v.emailPublisher.PublishVerificationToken(email, token)
	if err != nil {
		v.logger.WithError(err).Error("Failed to publish verification email")
	}

	v.logger.Info("Verification token generated and published")
	return nil
}

func (v *verificationUseCase) VerifyToken(token string) error {
	v.logger.WithField("token", token)

	data, err := v.verificationRepo.ValidateToken(token)
	if err != nil {
		v.logger.WithError(err).Warn("Invalid or expired token")
		return customErr.ErrVerificationTokenInvalid
	}

	err = v.userRepo.UpdateIsVerified(data.Email, true)
	if err != nil {
		v.logger.WithError(err).Error("Failed to update user verification status")
		return customErr.ErrInternalServer
	}

	err = v.verificationRepo.MarkTokenUsed(token)
	if err != nil {
		v.logger.WithError(err).Warn("Token used but failed to mark as used")
	}

	v.logger.WithField("email", data.Email).Info("User verified successfully")
	return nil
}

func (v *verificationUseCase) ResendVerification(email string) error {
	v.logger.WithField("email", email)

	user, err := v.userRepo.GetByEmail(email)
	if err != nil {
		v.logger.Warn("Resend verification failed: Email not found")
		return customErr.ErrLoginEmailNotFound
	}

	activeToken, err := v.verificationRepo.GetActiveVerificationByEmail(email)
	if err == nil && activeToken != nil {
		v.logger.Warn("Resend verification denied: Active token exists")
		return customErr.ErrVerificationTokenStillValid
	}

	token := uuid.NewString()
	expiresAt := time.Now().Add(15 * time.Minute)

	err = v.verificationRepo.CreateToken(user.Email, token, expiresAt)
	if err != nil {
		v.logger.WithError(err).Error("Failed to create verification token")
		return customErr.ErrInternalServer
	}

	err = v.emailPublisher.PublishVerificationToken(user.Email, token)
	if err != nil {
		v.logger.WithError(err).Error("Failed to publish verification email")
	}

	v.logger.Info("Verification token resent successfully")
	return nil
}
