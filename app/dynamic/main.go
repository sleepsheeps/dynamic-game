package main

import (
	"dynamic-game/config"
	"dynamic-game/dynamic"
	"dynamic-game/utils"
	"log"
)

func main() {
	// 初始化日志格式
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 初始化配置
	err := config.Config.LoadConfig()
	if err != nil {
		log.Fatal("load config error: ", err)
		return
	}
	// 初始化redis
	err = utils.InitCache(config.Config.Redis_ADDRESS)
	if err != nil {
		log.Fatal("init redis error", err)
		return
	}
	log.Println("cache init ok")
	// 初始化dynamic-server
	server := dynamic.GetIDynamic(config.Config)
	err = server.Init(config.Config)
	if err != nil {
		log.Println("server init error", err)
		return
	}
	log.Println(server.Name(), "dynamic server init ok")
	err = server.Run()
	if err != nil {
		log.Println("server run error", err)
		return
	}
	log.Println("server run ok")
}
