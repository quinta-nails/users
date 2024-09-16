package main

import (
	"buf.build/gen/go/quinta/nails/grpc/go/users/usersgrpc"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/quinta-nails/users/internal/db"
	"github.com/quinta-nails/users/internal/server"
	"google.golang.org/grpc"
	"log"
)

type Service struct {
	usersgrpc.UnimplementedUsersServiceServer
	db *db.Queries
}

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
	service := &Service{}

	dbConnection, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	service.db = dbConnection

	grpcServer := grpc.NewServer()
	usersgrpc.RegisterUsersServiceServer(grpcServer, service)

	err = server.ListenAndServe(grpcServer)
	if err != nil {
		log.Fatal(err)
	}
}
