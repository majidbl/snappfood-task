package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"task/models"
)

type Store struct {
	db *gorm.DB
}

func New() Store {
	return Store{
		db: NewDB(),
	}
}

type Fn func(context.Context, Store) error

func (s Store) Transaction(ctx context.Context, fn Fn) (err error) {
	tx := s.db.Begin()

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit().Error
		}
	}()

	newStore := Store{
		db: tx,
	}

	err = fn(ctx, newStore)
	return nil
}

func (s Store) GetOrderTrip(ctx context.Context, orderId uint) (models.Trip, error) {
	var t models.Trip
	result := s.db.First(&t, "order_id=?", orderId)
	return t, result.Error
}

func (s Store) CreateAgent(ctx context.Context, agent models.Agent) error {
	return s.db.Create(agent).Error
}
