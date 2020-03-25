package main

import (
	"dynamic-game/config"
	"dynamic-game/server"
	"log"
)

func main() {
	cfg := &config.Config{}
	cfg.Mode = "pool"
	s := &server.Server{}
	err := s.Start(cfg)
	if err != nil {
		log.Fatal("dynamic server start error: ", err)
		return
	}
	log.Println("dynamic server start ok")
}
