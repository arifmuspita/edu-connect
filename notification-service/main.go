package main

import (
	"notification_service/config"
	"notification_service/queue"
	"notification_service/repository"
	"notification_service/usecase"

	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	db := config.InitDB()

	logger := logrus.New()

	// migration.Migration(db)

	config.InitRabbitMQ()
	defer config.CloseRabbitMQ()

	notificationRepo := repository.NewNotificationRepository(db, logger)
	notificationUseCase := usecase.NewNotificationUsecase(notificationRepo, logger)

	go queue.StartConsumer(config.RabbitMQConn, notificationUseCase, logger)

	logger.Info("Notification Service is running...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Notification Service is running"))
	})

	port := os.Getenv("PORT")

	go func() {
		logger.Infof("Starting HTTP server on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			logger.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	select {}
}
