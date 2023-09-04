package mysqlstore

import (
	"context"

	"gorm.io/gorm"

	"task/models"
)

type IAgent interface {
	GetAgent(ctx context.Context, id int64) (models.Agent, error)
	UpdateAgent(ctx context.Context, agent *models.Agent) error
	CreateAgent(ctx context.Context, agent *models.Agent) error
}

type agent struct {
	db *gorm.DB
}

func NewAgent(db *gorm.DB) IAgent {
	return &agent{
		db: db,
	}
}

func (a agent) GetAgent(ctx context.Context, id int64) (models.Agent, error) {
	var agentModel models.Agent
	result := a.db.First(&agentModel, "id=?", id)
	return agentModel, result.Error
}

func (a agent) UpdateAgent(ctx context.Context, agent *models.Agent) error {
	return a.db.Updates(agent).Error
}

func (a agent) CreateAgent(ctx context.Context, agent *models.Agent) error {
	return a.db.Create(agent).Error
}
