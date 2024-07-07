package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitConfig() error {
	var err error

	db, err = InitSQLite()

	if err != nil {
		return fmt.Errorf("Erro ao iniciar o SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
