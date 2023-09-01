package models

import "gorm.io/gorm"

const (
	ReportReviewed   = "report_reviewed"
	ReportRegistered = "report_registered"
	ReportAssigned   = "report_assigned"
)

type DelayReport struct {
	*gorm.Model
	OrderID      uint   `json:"order_id"`
	AgentId      uint   `json:"agent_id"`
	DeliveryTime int    `json:"delivery_time"`
	Status       string `json:"status"`
}
