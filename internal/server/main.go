package server

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/quinta-nails/users/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func ListenAndServe(grpcServer *grpc.Server) error {
	cfg := config.ServerConfig{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	reflection.Register(grpcServer)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stopChan
		log.Println("Received stop signal, gracefully stopping the server...")
		grpcServer.GracefulStop()
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("server is listening on port %d...", cfg.Port))

	return grpcServer.Serve(listener)
}
