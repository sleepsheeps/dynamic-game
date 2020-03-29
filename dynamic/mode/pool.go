package mode

import (
	"dynamic-game/config"
	"dynamic-game/dynamic/pdynamic"
	"dynamic-game/proto/helper"
	"dynamic-game/utils"
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
	utils.RegisterMsg(helper.MsgType_T_Student, pdynamic.PTest)
	return nil
}

func (p *Pool) Run() {

}

func (p *Pool) Name() string {
	return p.name
}
