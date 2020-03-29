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

	// 初始化mq
	err = utils.StartMQ(config.Config.MQ)
	if err != nil {
		log.Fatal("start mq error", err)
		return
	}
	log.Println("mq start ok")
	// 初始化dynamic-server
	server := dynamic.GetIDynamic(config.Config)
	server.Init(config.Config)
	log.Println(server.Name(), "dynamic server start ok")
	startChan := make(chan struct{})
	<-startChan
}
