package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/TaskManager/config"
	tasks "github.com/TaskManager/internal"
)

var Db *gorm.DB

func NewConnectionPostgres() error {
	cfg := config.Parsed
	dsn := BuildDSNPostgres(
		cfg.TaskManagerPostgresHost,
		cfg.TaskManagerPostgresPort,
		cfg.TaskManagerPostgresUser,
		cfg.TaskManagerPostgresPassword,
		cfg.TaskManagerPostgresDB,
	)
	conn, errConn := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
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
	return nil
}

func BuildDSNPostgres(host, port, user, password, dbname string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func Migration() error {
	if errMigrate := Db.AutoMigrate(&tasks.MTask{}); errMigrate != nil {
		return errMigrate
	}
	return nil
}
