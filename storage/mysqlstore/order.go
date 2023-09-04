package mysqlstore

import (
	"context"

	"gorm.io/gorm"

	"task/models"
)

type IOrder interface {
	CreateOrder(ctx context.Context, order *models.Order) error
	UpdateOrder(ctx context.Context, order *models.Order) error
	GetOrderById(ctx context.Context, id uint) (models.Order, error)
	GetDelayedOrders(ctx context.Context) ([]models.Order, error)
	ExistOrder(ctx context.Context, id int64) (bool, error)
}

type order struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) IOrder {
	return &order{db: db}
}

func (o order) CreateOrder(ctx context.Context, order *models.Order) error {
	return o.db.Create(order).Error
}

func (o order) UpdateOrder(ctx context.Context, order *models.Order) error {
	return o.db.Updates(order).Error
}

func (o order) GetOrderById(ctx context.Context, id uint) (models.Order, error) {
	var order models.Order
	result := o.db.First(&order, "id=?", id)
	return order, result.Error
}

func (o order) GetDelayedOrders(ctx context.Context) ([]models.Order, error) {
	var order []models.Order
	result := o.db.Find(
		&order,
		"TIMESTAMPADD(MINUTE, delivery_time, created_at) < NOW() and state <> ?",
		models.OrderReview)
	return order, result.Error
}

func (o order) ExistOrder(ctx context.Context, id int64) (bool, error) {
	var exists bool
	var err error

	err = o.db.Model(&models.Order{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).
		Error

	return exists, err
}
