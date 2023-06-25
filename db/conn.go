package db

import (
	"covid/entity"
	"fmt"
	"log"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB
var SqlDB *sql.DB

func ConnectDB() {
	dbHost := "127.18.0.3"
	dbPort := "5432"
	dbUser := "root"
	dbName := "covid_db"
	dbPass := "root"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}
	s, err := db.DB()
	if err != nil {
		panic(err)
	}

	Conn = db
	SqlDB = s
}

func Migrate() {
	Conn.AutoMigrate(
		&entity.CovidSummary{},
	)
}
