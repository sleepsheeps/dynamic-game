package utils

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func StartMQ(addr string) error {
	if mq == nil {
		mq = new(MQ)
		mq.addr = addr
	}
	return mq.start()
}

var (
	mq *MQ
)

type MQ struct {
	addr string
	conn *nats.EncodedConn
}

func (mq *MQ) start() error {
	conn, err := nats.Connect(mq.addr)
	if err != nil {
		return err
	}
	c, _ := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	mq.conn = c
	return nil
}

func (mq *MQ) register(topic string) {
	mq.conn.Subscribe(topic, func(s ProtoMsg) {
		fmt.Println("!!!!", s.MsgType)
	})
}

func (mq *MQ) publish(topic string, msg ProtoMsg) error {
	return mq.conn.Publish(topic, msg)
}
