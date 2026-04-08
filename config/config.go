package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	looger *Logger
)

func InitDB() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("Initialize sqlite error: %v", err)
	}

	return nil
}

func GetSQLiteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize the logger
	logger := NewLogger(p)
	return logger
}
