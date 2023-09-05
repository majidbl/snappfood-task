package controllers

import "task/service"

type Controller struct {
	service service.INotificationService
}

func NewController(notificationService service.INotificationService) Controller {
	return Controller{service: notificationService}
}
