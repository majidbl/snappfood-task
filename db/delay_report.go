package db

import (
	"context"

	"task/models"
)

func (s Store) CreateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	return s.db.Create(delayReport).Error
}

func (s Store) UpdateDelayReport(ctx context.Context, delayReport *models.DelayReport) error {
	return s.db.Save(delayReport).Error
}

func (s Store) GetOrderDelayReport(ctx context.Context, orderId uint) (models.DelayReport, error) {
	var d models.DelayReport
	result := s.db.First(&d, "order_id=?", orderId)
	return d, result.Error
}
