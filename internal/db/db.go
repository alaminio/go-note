package db

import (
	"go-note/configs"
	"go-note/internal/note"
	"go-note/internal/todo"
	"log"

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

	// Auto migrate the schema
	err = DB.AutoMigrate(&note.Note{}, &todo.Todo{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected and migrated")
	return DB, nil
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
