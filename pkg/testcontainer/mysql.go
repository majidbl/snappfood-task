package testcontainer

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func StartMySQLContainer() (*gorm.DB, func(), error) {
	// Define a context for the container lifecycle.
	ctx := context.Background()

	// Create a MySQL container using testcontainers.
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0", // You can specify a specific MySQL version.
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",    // Change this to your desired root password.
			"MYSQL_ROOT_USER":     "root",    // Change this to your desired root password.
			"MYSQL_DATABASE":      "db_test", // Change this to your desired database name.
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlCtn, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return nil, nil, err
	}

	// Get the host port to connect to.
	hostPort, err := mysqlCtn.MappedPort(ctx, "3306")
	if err != nil {
		return nil, nil, fmt.Errorf("could not get mapped port: %v", err)
	}

	// Create a database connection.
	dsn := fmt.Sprintf("root:root@tcp(localhost:%s)/db_test", hostPort.Port())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("could not connect to MySQL: %v", err)
		//t.Fatalf("Could not connect to MySQL: %v", err)
	}

	// Wait for MySQL to be ready.
	//for i := 0; i < 30; i++ {
	//	pingErr := db.DB().Ping()
	//	if pingErr == nil {
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//}

	// Return a cleanup function to stop the container after the test.
	cleanup := func() {
		terminateErr := mysqlCtn.Terminate(ctx)
		if terminateErr != nil {
			log.Printf("Error stopping MySQL container: %v", terminateErr)
		}
	}

	if err != nil {
		return nil, nil, fmt.Errorf("MySQL did not become ready: %v", err)
	}

	return db, cleanup, nil
}
