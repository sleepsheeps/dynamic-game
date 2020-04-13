package allocate

import (
	"context"
	pdynamic "dynamic-game/proto/dynamic"
	"google.golang.org/grpc"
)

func register(s *grpc.Server) {
	pdynamic.RegisterAllocateServiceServer(s, new(AllocateServiceImpl))
}

type AllocateServiceImpl struct {
}

func (a *AllocateServiceImpl) Allocate(ctx context.Context, recv *pdynamic.GAllocateServer) (*pdynamic.AAllocateServer, error) {
	reply := &pdynamic.AAllocateServer{
		ServerID: []string{"test"},
	}
	return reply, nil
}
