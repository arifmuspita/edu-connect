package handler

import (
	"errors"
	"net/http"
	"userService/usecase"
	"userService/utils"

	customErr "userService/error"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type PasswordResetHandler struct {
	resetUC usecase.IPasswordResetUseCase
	logger  *logrus.Entry
}

func NewPasswordResetHandler(resetUC usecase.IPasswordResetUseCase, logger *logrus.Logger) *PasswordResetHandler {
	enrichedLogger := logger.WithField("layer", "handler")
	return &PasswordResetHandler{
		resetUC: resetUC,
		logger:  enrichedLogger,
	}
}

type RequestResetPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"new_password"`
}

// RequestResetPassword godoc
// @Summary Request password reset
// @Description Sends a password reset link to the user's email
// @Tags Password Reset
// @Accept json
// @Produce json
// @Param body body RequestResetPasswordRequest true "User email"
// @Success 200 {object} utils.APIResponse
// @Failure 400,500 {object} utils.APIResponse
// @Router /v1/forgot-password [post]
func (h *PasswordResetHandler) RequestResetPassword(c echo.Context) error {
	var req RequestResetPasswordRequest

	if err := c.Bind(&req); err != nil {
		h.logger.Warn("Invalid request body for RequestResetPassword")
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	h.logger.WithField("email", req.Email).Info("Password reset request received")

	err := h.resetUC.RequestReset(req.Email)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, customErr.ErrRegisterEmailRequired) || errors.Is(err, customErr.ErrLoginEmailNotFound) {
			statusCode = http.StatusBadRequest
		}
		h.logger.WithError(err).WithField("email", req.Email).Error("Password reset request failed")
		return utils.ErrorResponse(c, statusCode, err.Error())
	}

	h.logger.WithField("email", req.Email).Info("Password reset token sent successfully")

	return utils.SuccessResponse(c, http.StatusOK, nil, "Password reset link sent to email")
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Resets the user's password using the reset token
// @Tags Password Reset
// @Accept json
// @Produce json
// @Param token query string true "Reset Token"
// @Param body body ResetPasswordRequest true "New Password"
// @Success 200 {object} utils.APIResponse
// @Failure 400,500 {object} utils.APIResponse
// @Router /v1/reset-password [post]
func (h *PasswordResetHandler) ResetPassword(c echo.Context) error {

	token := c.QueryParam("token")

	var req ResetPasswordRequest

	if err := c.Bind(&req); err != nil {
		h.logger.Warn("Invalid request body for ResetPassword")
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	h.logger.WithField("token", token).Info("Reset password execution started")

	err := h.resetUC.ResetPassword(token, req.NewPassword)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, customErr.ErrVerificationTokenInvalid) || errors.Is(err, customErr.ErrRegisterInvalidPassword) {
			statusCode = http.StatusBadRequest
		}
		h.logger.WithError(err).WithField("token", token).Error("Reset password failed")
		return utils.ErrorResponse(c, statusCode, err.Error())
	}

	h.logger.WithField("token", token).Info("Password reset successfully")
	return utils.SuccessResponse(c, http.StatusOK, nil, "Password has been reset successfully")
}
