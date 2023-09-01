package models

import "gorm.io/gorm"

type Vendor struct {
	*gorm.Model
	Name string `json:"name"`
}

type VendorDelay struct {
	VendorID          uint
	VendorName        string
	OrderId           uint
	TotalDelayMinutes int
}
