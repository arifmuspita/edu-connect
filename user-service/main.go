package main

import (
	"os"
	"sync"
	"userService/config"
	"userService/grpc"
	"userService/handler"
	"userService/logger"
	"userService/middleware"
	"userService/queue"
	"userService/repository"
	"userService/route"
	"userService/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "userService/docs"
)

// @title User Service API
// @version 1.0
// @description This is a user service.
// @host localhost:8080
// @BasePath /v1
func main() {
	_ = godotenv.Load()

	logger.InitLogger()

	middleware.InitMiddlewareLogger(logger.Log)

	db := config.InitDB()

	rabbitConn, channel := config.InitRabbitMQ()
	defer rabbitConn.Close()
	defer channel.Close()

	// migration.Migration(db)

	emailPublisher, err := queue.NewEmailPublisher(channel, "email")
	if err != nil {
		panic("Failed to initialize email publisher: " + err.Error())
	}

	userRepo := repository.NewUserRepository(db, logger.Log)
	verificationRepo := repository.NewVerificationRepository(db, logger.Log)
	passwordResetRepo := repository.NewPasswordResetRepository(db, logger.Log)

	verificationUC := usecase.NewVerificationUseCase(userRepo, verificationRepo, emailPublisher, logger.Log)
	userUC := usecase.NewUserUseCase(userRepo, verificationUC, logger.Log)
	passwordResetUC := usecase.NewPasswordResetUseCase(userRepo, passwordResetRepo, emailPublisher, logger.Log)

	userHandler := handler.NewUserHandler(userUC, logger.Log)
	verificationHandler := handler.NewVerificationHandler(verificationUC, logger.Log)
	passwordResetHandler := handler.NewPasswordResetHandler(passwordResetUC, logger.Log)

	grpcPort := os.Getenv("GRPC_PORT")

	if grpcPort == "" {
		grpcPort = "50051"
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		grpc.StartGRPCServer(userRepo, grpcPort)
	}()

	go func() {
		defer wg.Done()

		e := echo.New()
		e.Use(middleware.LogrusMiddleware())
		route.Init(e, userHandler, *verificationHandler, *passwordResetHandler)
		e.GET("/swagger/*", echoSwagger.WrapHandler)

		e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	}()

	wg.Wait()

}
