//go:generate envdoc -format dotenv -output .env.example -dir . -files * -types *Config
package main

import (
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/quinta-nails/users/internal/db"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error loading .env file`)
	}

	err = db.ApplyMigrations()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	println(1)
}
