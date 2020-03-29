package mode

import (
	"dynamic-game/config"
	"dynamic-game/dynamic/pdynamic"
	"dynamic-game/proto/helper"
	"dynamic-game/utils"
)

type Allocate struct {
	name string
}

func (a *Allocate) Init(cfg *config.DynamicConfig) error {
	a.name = cfg.Mode
	utils.RegisterMsg(helper.MsgType_T_Student, pdynamic.PTest)
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
