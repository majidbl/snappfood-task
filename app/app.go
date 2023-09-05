package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"task/config"
	"task/controllers"
	"task/service"
	"task/storage/mysqlstore"
	"task/storage/queue"
)

func NewServer(config config.Config) {
	ctx := context.Background()
	queue.SetUpQueueManager(config)

	err := mysqlstore.SetUpDB(config)
	if err != nil {
		panic(err.Error())
	}

	store := mysqlstore.NewStore()
	vendor := mysqlstore.NewVendor(mysqlstore.NewDB())

	notificationService := service.NewService(store, vendor)
	ctrl := controllers.NewController(notificationService)

	start, stop := NewEchoServer(ctx, config, ctrl)

	go func() {
		if startServerErr := start(); startServerErr != nil {
			panic(startServerErr)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		logrus.Errorf("signal.Notify: %v", v)
		stopHttpErr := stop()
		if stopHttpErr != nil {
			logrus.Errorf("stopHttpErr: %v", stopHttpErr)
		}
	case done := <-ctx.Done():
		logrus.Errorf("ctx.Done: %v", done)
	}
}
