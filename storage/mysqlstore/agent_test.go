package mysqlstore

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"testing"

	"github.com/stretchr/testify/require"

	"task/models"
)

func TestCreateOneAgent(t *testing.T) {
	ctx := context.Background()
	err := agentTest.CreateAgent(ctx, &models.Agent{
		Name:   "test-agent-1",
		Status: "",
	})

	require.NoError(t, err)
}

func TestCreateMultipleAgent(t *testing.T) {
	testCases := []struct {
		name  string
		agent *models.Agent
		code  [5]byte
	}{
		{
			name: "test agent 1",
			agent: &models.Agent{
				Name:   "test agent 1",
				Status: models.Busy,
			},
			code: [5]byte{},
		},
		{
			name: "test agent 1",
			agent: &models.Agent{
				Name:   "test agent 1",
				Status: models.Busy,
			},
			code: [5]byte{'2', '3', '0', '0', '0'},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.Background()
			err := agentTest.CreateAgent(ctx, test.agent)
			if err != nil {
				var e *mysql.MySQLError
				errors.As(err, &e)
				if e.SQLState != test.code {
					t.Errorf("got err %s, want %s", e.SQLState, test.code)
				}
			}
		})
	}
}
