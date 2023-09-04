package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"task/config"
	"task/controllers"
	"task/docs"
	"task/service"
	"task/storage/mysqlstore"
	"task/storage/queue"
)

func NewServer(config config.Config) {
	queue.SetUpQueueManager(config)

	err := mysqlstore.SetUpDB(config)
	if err != nil {
		panic(err.Error())
	}

	store := mysqlstore.NewStore()
	vendor := mysqlstore.NewVendor(mysqlstore.NewDB())

	notificationService := service.NewService(store, vendor)
	ctrl := controllers.NewController(notificationService)

	e := echo.New()

	docs.SwaggerInfo.Title = "Delivery Notification REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("api/v1")
	v1.POST("/delay/report", ctrl.ReportDelay())
	v1.POST("/delay/assign", ctrl.AssignDelay())
	v1.GET("/delay/report", ctrl.DelayReport())

	e.Logger.Fatal(e.Start(config.HTTPServerAddress))
}
