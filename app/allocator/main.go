package main

import (
	"dynamic-game/allocator"
	"log"
)

func main() {
	err := allocator.Start()
	if err != nil {
		log.Fatal("allocator start error", err)
	}
}
