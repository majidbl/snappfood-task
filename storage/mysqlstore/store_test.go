package mysqlstore

import (
	"log"
	"os"
	"testing"

	"gorm.io/gorm"

	"task/pkg/testcontainer"
)

var (
	testDB    *gorm.DB
	agentTest IAgent
)

func TestMain(m *testing.M) {
	var err error
	var terminator func()

	testDB, terminator, err = testcontainer.StartMySQLContainer()
	if err != nil {
		log.Fatal(err)
	}

	defer terminator()
	err = MigrateUp(testDB)
	if err != nil {
		log.Fatal(err)
	}

	agentTest = NewAgent(testDB)

	os.Exit(m.Run())
}
