package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	looger *Logger
)

func InitDB() error {
	return nil
}


func GetLogger(p string) *Logger {
	// Initialize the logger
	logger := NewLogger(p)
	return logger
}