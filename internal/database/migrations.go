package database

import (
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	"github.com/caarlos0/env/v11"
	"github.com/quinta-nails/users/internal/config"
)

func ApplyMigrations() error {
	cfg := config.Database{}
	// Валидация переменных окружения
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	db := dbmate.New(&cfg.URL)

	err := db.CreateAndMigrate()
	if err != nil {
		return err
	}

	return nil
}
