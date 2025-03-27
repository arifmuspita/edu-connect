package handler

import (
	"net/http"
	"userService/usecase"
	"userService/utils"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	customErr "userService/error"
)

type requestResendVerify struct {
	Email string `json:"email"`
}

type VerificationHandler struct {
	verificationUC usecase.IVerificationUseCase
	logger         *logrus.Entry
}

func NewVerificationHandler(verificationUC usecase.IVerificationUseCase, logger *logrus.Logger) *VerificationHandler {
	enrichedLogger := logger.WithField("layer", "handler")
	return &VerificationHandler{
		verificationUC: verificationUC,
		logger:         enrichedLogger,
	}
}

// Verify godoc
// @Summary Verify email with token
// @Description Verifies the user's email using a token sent via email
// @Tags Verification
// @Accept json
// @Produce json
// @Param token query string true "Verification Token"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /v1/verify [get]
func (h *VerificationHandler) Verify(c echo.Context) error {

	token := c.QueryParam("token")

	if token == "" {
		h.logger.Warn("Verification failed: token is missing")
		return utils.ErrorResponse(c, http.StatusBadRequest, "Token is required")
	}

	h.logger.WithField("token", token).Info("Verification request received")

	err := h.verificationUC.VerifyToken(token)
	if err != nil {
		if err == customErr.ErrVerificationTokenInvalid {
			h.logger.WithField("token", token).Warn("Verification failed: invalid or expired token")
			return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid or expired token")
		}

		h.logger.WithField("token", token).Error("Verification failed: internal error")
		return utils.ErrorResponse(c, http.StatusBadRequest, "Internal server error")
	}

	h.logger.WithField("token", token).Info("User verified successfully")
	return utils.SuccessResponse(c, http.StatusOK, nil, "Email verified successfully")

}

// ResendVerification godoc
// @Summary Resend email verification
// @Description Sends a new verification email if no valid token exists
// @Tags Verification
// @Accept json
// @Produce json
// @Param body body requestResendVerify true "Email for Resend Verification"
// @Success 200 {object} utils.APIResponse
// @Failure 400,500 {object} utils.APIResponse
// @Router /v1/resend-verification [post]
func (h *VerificationHandler) ResendVerification(c echo.Context) error {
	var req requestResendVerify
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	if req.Email == "" {
		h.logger.Warn("Resend verification failed: email is missing")
		return utils.ErrorResponse(c, http.StatusBadRequest, "Email is required")
	}

	h.logger.WithField("email", req.Email).Info("Resend verification request received")

	err := h.verificationUC.ResendVerification(req.Email)
	if err != nil {
		if err == customErr.ErrLoginEmailNotFound {
			h.logger.WithField("email", req.Email).Warn("Resend verification failed: Email not found")
			return utils.ErrorResponse(c, http.StatusBadRequest, "Email not found")
		}
		if err == customErr.ErrVerificationTokenStillValid {
			h.logger.WithField("email", req.Email).Warn("Resend verification denied: Active token exists")
			return utils.ErrorResponse(c, http.StatusBadRequest, "Existing verification token still valid")
		}

		h.logger.WithField("email", req.Email).Error("Resend verification failed: internal error")
		return utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
	}

	h.logger.WithField("email", req.Email).Info("Verification email resent successfully")
	return utils.SuccessResponse(c, http.StatusOK, nil, "Verification email resent successfully")
}
