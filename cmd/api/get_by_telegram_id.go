package main

import (
	pb "buf.build/gen/go/quinta/nails/protocolbuffers/go/users"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GetByTelegramID retrieves a user by their Telegram ID.
// It takes a context.Context and a pb.GetByTelegramIDRequest as input,
// and returns a pb.GetByTelegramIDResponse and an error as output.
func (s *Service) GetByTelegramID(ctx context.Context, in *pb.GetByTelegramIDRequest) (*pb.GetByTelegramIDResponse, error) {
	return &pb.GetByTelegramIDResponse{Result: &pb.User{
		Id:         1,
		TelegramId: 1,
		FirstName:  "dasd",
		CreatedAt:  timestamppb.Now(),
	}}, nil
}
