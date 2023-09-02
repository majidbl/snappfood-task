package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"task/controllers"
	"task/docs"
	"task/storage/mysql"
	"task/storage/queue"
	"task/util"
)

func NewServer(config util.Config) {
	queue.SetUpQueueManager(config)

	err := mysql.SetUpDB(config)
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	docs.SwaggerInfo.Title = "Delivery Notification REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("api/v1")
	v1.POST("/delay/report", controllers.ReportDelay())
	v1.POST("/delay/assign", controllers.AssignDelay())
	v1.GET("/delay/report", controllers.DelayReport())

	e.Logger.Fatal(e.Start(config.HTTPServerAddress))
}
