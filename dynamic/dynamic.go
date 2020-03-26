package dynamic

import (
	"dynamic-game/config"
	"dynamic-game/mode"
)

func init() {
	registerMgr()
}

var (
	dynamicMgr map[string]IDynamic
)

type IDynamic interface {
	Init(cfg *config.Config) error
	Name() string
	Run()
}

func registerMgr() {
	dynamicMgr = make(map[string]IDynamic)
	dynamicMgr[config.MODE_ALLOCATE] = mode.NewAllocate()
	dynamicMgr[config.MODE_POOL] = mode.NewPool()
}

func GetIDynamic(config2 *config.Config) IDynamic {
	return dynamicMgr[config2.Mode]
}
