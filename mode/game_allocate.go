package mode

import (
	"dynamic-game/config"
	"log"
)

type Allocate struct {
}

func (a *Allocate) Init(cfg *config.Config) error {
	log.Println("allocate init ok")
	return nil
}

func (a *Allocate) Run() {

}

func NewAllocate() *Allocate {
	a := &Allocate{}
	return a
}
