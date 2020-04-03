package allocate

import (
	"dynamic-game/config"
)

func NewAllocate() *Allocate {
	a := &Allocate{}
	return a
}

type Allocate struct {
	name string
}

func (a *Allocate) Init(cfg *config.DynamicConfig) error {
	a.name = cfg.Mode
	return nil
}

func (a *Allocate) Run() {

}

func (a *Allocate) Name() string {
	return a.name
}

func (a *Allocate) Monitor() {

}
