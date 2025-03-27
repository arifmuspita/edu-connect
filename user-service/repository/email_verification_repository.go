package repository

import (
	"time"
	"userService/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IVerificationRepository interface {
	CreateToken(email, token string, expiresAt time.Time) error
	ValidateToken(token string) (*model.EmailVerification, error)
	MarkTokenUsed(token string) error
	GetActiveVerificationByEmail(email string) (*model.EmailVerification, error)
}

type verificationRepository struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func NewVerificationRepository(db *gorm.DB, baseLogger *logrus.Logger) IVerificationRepository {
	enrichedLogger := baseLogger.WithField("layer", "repository")
	return &verificationRepository{
		db:     db,
		logger: enrichedLogger,
	}
}

func (r *verificationRepository) CreateToken(email, token string, expiresAt time.Time) error {
	data := &model.EmailVerification{
		Email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
		Used:      false,
	}
	err := r.db.Create(data).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"email": email,
			"token": token,
			"error": err,
		}).Error("Failed to create email verification token")
		return err
	}

	r.logger.WithFields(logrus.Fields{
		"email": email,
		"token": token,
	}).Info("Verification token created successfully")

	return nil
}

func (r *verificationRepository) ValidateToken(token string) (*model.EmailVerification, error) {
	var ev model.EmailVerification
	err := r.db.Where("token = ? AND used = false AND expires_at > ?", token, time.Now()).First(&ev).Error

	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"token": token,
			"error": err,
		}).Warn("Failed to validate token (might be expired or invalid)")
		return nil, err
	}

	r.logger.WithField("token", token).Info("Verification token validated")
	return &ev, nil
}

func (r *verificationRepository) MarkTokenUsed(token string) error {
	err := r.db.Model(&model.EmailVerification{}).
		Where("token = ?", token).
		Update("used", true).Error

	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"token": token,
			"error": err,
		}).Warn("Failed to mark verification token as used")
		return err
	}

	r.logger.WithField("token", token).Info("Verification token marked as used")
	return nil
}

func (r *verificationRepository) GetActiveVerificationByEmail(email string) (*model.EmailVerification, error) {
	var ev model.EmailVerification
	err := r.db.Where("email = ? AND used = false AND expires_at > ?", email, time.Now()).First(&ev).Error

	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"email": email,
			"error": err,
		}).Warn("No active verification token found")
		return nil, err
	}

	r.logger.WithField("email", email).Info("Active verification token found")
	return &ev, nil
}
