package models

import "gorm.io/gorm"

const (
	OrderRegistered = "order_registered"
	OrderReview     = "order_review"
	OrderDelivered  = "order_delivered"
)

type Order struct {
	*gorm.Model
	VendorId     int64  `json:"vendor_id"`
	DeliveryTime int64  `json:"delivery_time"`
	State        string `json:"state"`
}
