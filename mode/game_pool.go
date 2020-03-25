package mode

import (
	"dynamic-game/config"
	"log"
)

type Pool struct {
}

func NewPool() *Pool {
	p := &Pool{}
	return p
}

func (p *Pool) Init(cfg *config.Config) error {
	log.Println("pool init ok")
	return nil
}

func (p *Pool) Run() {

}
