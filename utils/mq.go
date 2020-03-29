package utils

import (
	"dynamic-game/config"
	"fmt"
	"github.com/nats-io/nats.go"
)

func StartMQ(addr string) error {
	if mq == nil {
		mq = new(MQ)
		mq.addr = addr
		mq.ch = make(chan *ProtoMsg)
	}
	return mq.start()
}

var (
	mq *MQ
)

type MQ struct {
	addr string
	conn *nats.EncodedConn
	ch   chan *ProtoMsg
}

func (mq *MQ) start() error {
	conn, err := nats.Connect(mq.addr)
	if err != nil {
		return err
	}
	c, _ := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	mq.conn = c
	mq.recv(config.Config.ServerID)
	return nil
}

func (mq *MQ) recv(topic string) {
	mq.conn.BindRecvChan(topic, mq.ch)
	go func() {
		select {
		case recv := <-mq.ch:
			fmt.Println(recv)
		}
	}()
}

func (mq *MQ) publish(topic string, msg *ProtoMsg) error {
	return mq.conn.Publish(topic, msg)
}
