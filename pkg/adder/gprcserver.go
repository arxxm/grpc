package adder

import (
	"context"

	api "github.com/arxxm/grpc/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

}
