package mysql2

import (
	"context"
	"gorm.io/gorm"

	"task/models"
)

type IDelayReport interface {
	CreateDelayReport(ctx context.Context, delayReport *models.DelayReport) error
	UpdateDelayReport(ctx context.Context, delayReport *models.DelayReport) error
	GetOrderDelayReport(ctx context.Context, orderId uint) (models.DelayReport, error)
}

type delayReport struct {
	db *gorm.DB
}

func NewDelayReport(db *gorm.DB) IDelayReport {
	return &delayReport{db: db}
}

func (d delayReport) CreateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	return d.db.Create(delayReport).Error
}

func (d delayReport) UpdateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	return d.db.Updates(delayReport).Error
}

func (d delayReport) GetOrderDelayReport(ctx context.Context, orderId uint) (models.DelayReport, error) {
	var delay models.DelayReport
	result := d.db.First(&delay, "order_id=?", orderId)
	return delay, result.Error
}
