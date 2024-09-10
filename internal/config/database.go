package config

import "net/url"

type DatabaseConfig struct {
	URL url.URL `env:"DATABASE_UR,required" envDefault:"postgres://postgres:postgres@localhost:5432/users?sslmode=disable"`
}
