package config

import "errors"

const (
	MODE_ALLOCATE = "ALLOCATE"
	MODE_POOL     = "POOL"
)

func init() {
	Config = &DynamicConfig{}
}

type DynamicConfig struct {
	Replicas       int
	Limit          int
	Mode           string
	Redis_ADDRESS  string
	Redis_PWD      string
	FleetName      string
	NameSpace      string
	Allocator_Addr string
}

var Config *DynamicConfig

func (c *DynamicConfig) LoadConfig() error {
	if c == nil {
		return errors.New("config struct is nil")
	}
	c.Mode = "POOL"
	c.Redis_ADDRESS = ":6379"
	return nil
}
