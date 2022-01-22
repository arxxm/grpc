package main

import (
	"log"
	"net"

	"github.com/arxxm/grpc/blob/master/pkg/api/adder.pb.go/pkg/adder"
	gen "github.com/arxxm/grpc/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	gen.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
