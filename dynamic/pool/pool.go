package pool

import (
	"dynamic-game/config"
)

type Pool struct {
	name string
}

func NewPool() *Pool {
	p := &Pool{}
	return p
}

func (p *Pool) Init(cfg *config.DynamicConfig) error {
	p.name = cfg.Mode
	return nil
}

func (p *Pool) Run() error {
	return nil
}

func (p *Pool) Name() string {
	return p.name
}
