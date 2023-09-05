package app

import (
	"context"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"task/config"
	"task/controllers"
	"task/docs"
)

type Start func() error
type Shutdown func() error

func NewEchoServer(ctx context.Context, config config.Config, ctrl controllers.Controller) (Start, Shutdown) {
	e := echo.New()

	docs.SwaggerInfo.Title = "Delivery Notification REST API"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("api/v1")
	v1.POST("/delay/report", ctrl.ReportDelay())
	v1.POST("/delay/assign", ctrl.AssignDelay())
	v1.GET("/delay/report", ctrl.DelayReport())

	start := func() error {
		e.Logger.Fatal(e.Start(config.HTTPServerAddress))
		return nil
	}

	shutdown := func() error {
		return e.Shutdown(ctx)
	}

	return start, shutdown
}
