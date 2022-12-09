package validator

import (
	"context"
	"google.golang.org/grpc"
	"grcpValidatorIPv4/api"
	"log"
	"net"
)

type GRPCServer struct {
	api.UnimplementedValidatorServer
}

func (s *GRPCServer) Validate(ctx context.Context, request *api.ValRequest) (*api.ValidatedResponse, error) {
	if v := net.ParseIP(request.Ipv4); v == nil {
		return &api.ValidatedResponse{Response: false}, nil
	}

	return &api.ValidatedResponse{Response: true}, nil
}

func ListenAndServe(srv *GRPCServer, errCh chan error) {
	s := grpc.NewServer()
	api.RegisterValidatorServer(s, srv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		errCh <- err
		log.Fatal(err)
	}

	log.Println("[DEBUG] Running grcp server on port 8080")
	if err = s.Serve(l); err != nil {
		errCh <- err
		log.Fatal(err)
	}
}
