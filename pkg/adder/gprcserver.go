package adder

import (
	"context"
)

// GRPCServer ...
type GRPCServer struct{}

func (s *GRPCServer) Add(ctx context.Context, req *gen.AddRequest) (*gen.AddResponse, error) {

}
