package dynamic

import (
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	c, err := Client()
	if err != nil {
		log.Println("get client error", err)
		return
	}
	err = c.Allocate()
	if err != nil {
		log.Println(err)
	}
}
