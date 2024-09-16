package db

import (
	"database/sql"
	"github.com/caarlos0/env/v11"
	"github.com/quinta-nails/users/internal/config"
)

func NewDB() (*Queries, error) {
	cfg := config.DatabaseConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", cfg.URL.String())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return New(db), nil

}
