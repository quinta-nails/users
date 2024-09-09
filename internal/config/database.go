package config

import "net/url"

type Database struct {
	URL url.URL `env:"DATABASE_URL,required"`
}
