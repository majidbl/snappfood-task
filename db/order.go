package db

import (
	"context"
	"task/models"
)

func (s Store) CreateOrder(ctx context.Context, order *models.Order) error {
	return s.db.Create(order).Error
}

func (s Store) UpdateOrder(ctx context.Context, order *models.Order) error {
	return s.db.Updates(order).Error
}

func (s Store) GetDelayedOrder(ctx context.Context, id uint) (models.Order, error) {
	var order models.Order
	result := s.db.First(&order, "id=?", id)
	return order, result.Error
}

func (s Store) GetDelayedOrders(ctx context.Context) ([]models.Order, error) {
	var order []models.Order
	result := s.db.Find(
		&order,
		"TIMESTAMPADD(MINUTE, delivery_time, created_at) < NOW() and state <> ?",
		models.OrderReview)
	return order, result.Error
}

func (s Store) ExistOrder(ctx context.Context, id int64) (bool, error) {
	var exists bool
	var err error

	err = s.db.Model(&models.Order{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error

	return exists, err
}
