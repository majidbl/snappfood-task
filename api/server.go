package api

import (
	"sync"
	"task/db"
	"task/pkg/queue"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"task/docs"
	"task/models"
	"task/util"
)

var OrderQueueManger queue.Queue[models.Order]
var once sync.Once

func NewServer(config util.Config) {
	setUpQueueManager(config)
	err := db.SetUpDB(config)
	if err != nil {
		panic(err.Error())
	}

	e := echo.New()

	docs.SwaggerInfo.Title = "Promissory REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("api/v1")
	v1.POST("/delay/report", ReportDelay())
	v1.POST("/delay/assign", AssignDelay())
	v1.GET("/delay/report", DelayReport())

	e.Logger.Fatal(e.Start(config.HTTPServerAddress))
}

func setUpQueueManager(config util.Config) {
	once.Do(func() {
		OrderQueueManger = queue.New[models.Order](config)
	})
}
