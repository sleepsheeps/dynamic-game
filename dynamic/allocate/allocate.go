package allocate

import (
	"dynamic-game/config"
	"google.golang.org/grpc"
	"net"
)

func NewAllocate() *Allocate {
	a := &Allocate{}
	return a
}

type Allocate struct {
	name string
	*grpc.Server
	serverAddr string
}

func (a *Allocate) Init(cfg *config.DynamicConfig) error {
	a.name = cfg.Mode
	a.Server = grpc.NewServer()
	a.serverAddr = cfg.ServerAddr
	return nil
}

func (a *Allocate) Run() error {
	register(a.Server)
	lis, err := net.Listen("tcp", a.serverAddr)
	if err != nil {
		return err
	}
	err = a.Serve(lis)
	return err
}

func (a *Allocate) Name() string {
	return a.name
}
