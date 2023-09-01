package models

import "gorm.io/gorm"

const (
	DELIVERED = "delivered"
	PICKED    = "picked"
	AtVENDOR  = "at_vendor"
	ASSIGNED  = "assigned"
)

type Trip struct {
	*gorm.Model
	OrderID int    `json:"order_id"`
	Status  string `json:"status"` // Possible values: "ASSIGNED", "AT_VENDOR", "PICKED", "DELIVERED"
}
