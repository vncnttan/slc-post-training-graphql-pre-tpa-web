package database

import (
	"github.com/vncnttan/TrainingGraphQL/graph/model"
	"github.com/vncnttan/TrainingGraphQL/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

const defaultDatabase = "host=localhost user=gorm password=123456 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func GetInstance() *gorm.DB {
	if database == nil {
		dsn := helper.GoDotEnvVariable("DATABASE_URL")

		if dsn == "" {
			dsn = defaultDatabase
		}

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		database = db
	}

	return database
}

func MigrateTable() {
	db := GetInstance()
	db.AutoMigrate(&model.User{}, &model.Tweet{})
}
