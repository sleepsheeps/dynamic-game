package server

import (
	"dynamic-game/config"
	"dynamic-game/dynamic"
)

type Server struct {
	dynamic.IDynamic
}

func (s *Server) Start(config *config.Config) error {
	s.IDynamic = dynamic.GetIDynamic(config)
	err := s.Init(config)
	return err
}
