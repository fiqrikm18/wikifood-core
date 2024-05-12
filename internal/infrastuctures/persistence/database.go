package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type DBConnection struct {
	DB *gorm.DB
}

func NewDBConnection() *DBConnection {
	dbMaxIdleConnTime, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN_TIME"))
	dbMaxLifeTime, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN_TIME"))

	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbTimeZone := os.Getenv("DB_TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
		dbTimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	conn.SetMaxIdleConns(dbMaxIdleConnTime)
	conn.SetConnMaxIdleTime(time.Minute) // 1 minute idle time
	conn.SetMaxIdleConns(dbMaxLifeTime)
	conn.SetConnMaxLifetime(time.Minute) // 1 minute idle time

	return &DBConnection{
		DB: db,
	}
}
