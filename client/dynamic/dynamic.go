package dynamic

import (
	"context"
	"dynamic-game/config"
	pdynamic "dynamic-game/proto/dynamic"
	"google.golang.org/grpc"
	"log"
)

type client struct {
	*grpc.ClientConn
}

func (c *client) Allocate() error {
	allocateClient := pdynamic.NewAllocateServiceClient(c.ClientConn)
	reply, err := allocateClient.Allocate(context.Background(), &pdynamic.GAllocateServer{
		RequestID: []int32{1, 2, 3},
	})
	if err != nil {
		log.Println("allocate error", err)
		return err
	}
	log.Println("receive allocate:", reply.ServerID)
	return nil
}

func (c *client) Close() error {
	err := c.ClientConn.Close()
	return err
}

var c *client

func Client() (*client, error) {
	if c == nil {
		c = &client{}
		err := config.Config.LoadConfig()
		if err != nil {
			return nil, err
		}
		c.ClientConn, err = grpc.Dial(config.Config.ServerAddr, grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}
