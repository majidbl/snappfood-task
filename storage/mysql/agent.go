package mysql

import (
	"context"

	"task/models"
)

func (s Store) GetAgent(ctx context.Context, id int64) (models.Agent, error) {
	var agent models.Agent
	result := s.db.First(&agent, "id=?", id)
	return agent, result.Error
}

func (s Store) UpdateAgent(ctx context.Context, agent *models.Agent) error {
	return s.db.Updates(agent).Error
}
