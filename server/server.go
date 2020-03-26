package server

import (
	"dynamic-game/config"
	"dynamic-game/dynamic"
	"dynamic-game/utils"
	"log"
)

type Server struct {
	dynamic.IDynamic
}

func (s *Server) Start(config *config.DynamicConfig) error {
	// 初始化日志格式
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 初始化redis
	err := utils.InitCache(config.Redis_ADDRESS)
	if err != nil {
		return err
	}
	log.Println("cache init ok")
	s.IDynamic = dynamic.GetIDynamic(config)
	err = s.Init(config)
	return err
}
