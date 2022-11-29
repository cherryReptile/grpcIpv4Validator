package validator

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grcpValidatorIPv4/api"
	"log"
)

func ClientDo(request *api.ValRequest) (*api.ValidatedResponse, error) {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c := api.NewValidatorClient(conn)
	res, err := c.Validate(context.Background(), request)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}
