package database

import (
  app "app"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := app.GetConfig("DB_HOST")
	user := app.GetConfig("DB_USER")
	password := app.GetConfig("DB_PASSWORD")
	dbname := app.GetConfig("DB_NAME")
	port := app.GetConfig("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connection Opened to Database")

  DB = database
}
