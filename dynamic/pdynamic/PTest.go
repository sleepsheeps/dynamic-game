package pdynamic

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func PTest(message proto.Message) error {
	fmt.Println(message.String())
	fmt.Println(message.ProtoMessage)
	return nil
}
