package main

import (
	"dynamic-game/config"
	"dynamic-game/server"
	"log"
)

func main() {
	cfg := &config.Config{}
	cfg.Mode = "POOL"
	s := &server.Server{}
	err := s.Start(cfg)
	if err != nil {
		log.Fatal("dynamic server start error: ", err)
		return
	}
	log.Println(s.Name(), "dynamic server start ok")
}
