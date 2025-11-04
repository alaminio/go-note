package db

import (
	"database/sql"
	"go-note/configs"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	config := configs.GetConfig()
	var err error

	switch config.DBDriver {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	default:
		DB, err = gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	// Run migrations
	err = RunMigrations(config.DBName)
	if err != nil {
		return nil, err
	}

	log.Println("Database connected and migrated")
	return DB, nil
}

func RunMigrations(dbPath string) error {
	// Open database connection for migration
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	// Create migrate instance
	m, err := migrate.New(
		"file://migrations",
		"sqlite3://"+dbPath,
	)
	if err != nil {
		return err
	}
	defer m.Close()

	// Run all up migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Failed to get underlying sql.DB:", err)
		return
	}
	sqlDB.Close()
	log.Println("Database connection closed")
}
