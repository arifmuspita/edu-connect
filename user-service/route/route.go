package route

import (
	"userService/handler"
	"userService/middleware"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo,
	userHandler handler.UserHandler,
	verificationHandler handler.VerificationHandler,
	passwordResetHandler handler.PasswordResetHandler) {

	v1 := e.Group("/v1")
	v1.Use(middleware.LogrusMiddleware())

	v1.POST("/register", userHandler.Register)

	v1.POST("/login", userHandler.Login)

	v1.GET("/verify", verificationHandler.Verify)

	v1.POST("/forgot-password", passwordResetHandler.RequestResetPassword)

	v1.POST("/reset-password", passwordResetHandler.ResetPassword)

	v1.POST("/resend-verification", verificationHandler.ResendVerification)

	user := v1.Group("/users")

	user.GET("/:id", userHandler.GetUserByID)

	user.GET("", userHandler.GetAllUsersPaginated)

}
