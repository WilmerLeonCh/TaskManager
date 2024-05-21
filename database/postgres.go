package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func NewConnectionPostgres(dsn string) error {
	conn, errConn := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errConn != nil {
		return fmt.Errorf("failed to open connection to database: %w", errConn)
	}
	postgresDB, errDB := conn.DB()
	if errDB != nil {
		return fmt.Errorf("failed to get generic database object: %w", errDB)
	}
	if errPing := postgresDB.Ping(); errPing != nil {
		return fmt.Errorf("failed to ping database: %w", errPing)
	}
	Db = conn
	log.Println("Connected to database successfully!")
	return nil
}
