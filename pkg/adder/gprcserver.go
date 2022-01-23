package adder

import (
	"context"

	api "github.com/arxxm/grpc/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
