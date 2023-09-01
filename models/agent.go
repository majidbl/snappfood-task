package models

import "gorm.io/gorm"

const (
	Free = "free"
	Busy = "busy"
)

type Agent struct {
	*gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}
