package mysqlstore

import (
	"sync"
	"task/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func NewDB() *gorm.DB {
	return DB
}

func SetUpDB(config config.Config) error {
	var dbErr error
	once.Do(func() {
		mysqlDB, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
		if err != nil {
			dbErr = err
		}

		DB = mysqlDB
	})

	return dbErr
}
