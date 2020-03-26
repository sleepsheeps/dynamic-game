package mode

import (
	"dynamic-game/config"
	"log"
)

type Allocate struct {
	name string
}

func (a *Allocate) Init(cfg *config.Config) error {
	log.Println("allocate init ok")
	a.name = cfg.Mode
	return nil
}

func (a *Allocate) Run() {

}

func (a *Allocate) Name() string {
	return a.name
}

func NewAllocate() *Allocate {
	a := &Allocate{}
	return a
}
