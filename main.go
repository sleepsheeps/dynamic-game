package main

import (
	"dynamic-game/config"
	"dynamic-game/server"
	"log"
)

func main() {
	err := config.Config.LoadConfig()
	if err != nil {
		log.Fatal("load config error: ", err)
		return
	}
	s := &server.Server{}
	err = s.Start(config.Config)
	if err != nil {
		log.Fatal("dynamic server start error: ", err)
		return
	}
	log.Println(s.Name(), "dynamic server start ok")
}
