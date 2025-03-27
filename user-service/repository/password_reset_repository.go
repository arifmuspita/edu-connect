package repository

import (
	"time"
	"userService/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IPasswordResetRepository interface {
	CreateResetToken(email, token string, expiresAt time.Time) error
	ValidateResetToken(token string) (*model.PasswordReset, error)
	MarkResetTokenUsed(token string) error
	GetActivePasswordResetByEmail(email string) (*model.PasswordReset, error)
}

type passwordResetRepository struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func NewPasswordResetRepository(db *gorm.DB, baseLogger *logrus.Logger) IPasswordResetRepository {
	enrichedLogger := baseLogger.WithField("layer", "repository")
	return &passwordResetRepository{
		db:     db,
		logger: enrichedLogger,
	}
}

func (r *passwordResetRepository) CreateResetToken(email, token string, expiresAt time.Time) error {

	data := &model.PasswordReset{
		Email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
		Used:      false,
	}

	err := r.db.Create(data).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Error("Failed to create reset token")
		return err
	}

	r.logger.WithField("email", email).Info("Reset token created successfully")
	return nil
}

func (r *passwordResetRepository) ValidateResetToken(token string) (*model.PasswordReset, error) {
	var reset model.PasswordReset
	err := r.db.Where("token = ? AND used = false AND expires_at > ?", token, time.Now()).
		Limit(1).
		First(&reset).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"token": token,
			"error": err,
		}).Warn("Invalid or expired reset token")
		return nil, err
	}
	r.logger.WithField("token", token).Info("Reset token validated")
	return &reset, nil
}

func (r *passwordResetRepository) MarkResetTokenUsed(token string) error {
	err := r.db.Model(&model.PasswordReset{}).
		Where("token = ?", token).
		Update("used", true).Error

	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"token": token,
			"error": err,
		}).Warn("Failed to mark reset token as used")
		return err
	}

	r.logger.WithField("token", token).Info("Reset token marked as used")
	return nil
}

func (r *passwordResetRepository) GetActivePasswordResetByEmail(email string) (*model.PasswordReset, error) {
	var reset model.PasswordReset
	err := r.db.Where("email = ? AND used = ? AND expires_at > ?", email, false, time.Now()).
		First(&reset).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.logger.WithField("email", email).Info("No active reset token found")
		} else {
			r.logger.WithFields(logrus.Fields{
				"email": email,
				"error": err,
			}).Error("Failed to get active password reset token")
		}
		return nil, err
	}

	r.logger.WithField("email", email).Info("Active reset token found")
	return &reset, nil
}
