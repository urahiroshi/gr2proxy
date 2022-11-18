package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

const grpcHeaderLength = 5

func toJson(b []byte, messageDesc *desc.MessageDescriptor) ([]byte, error) {
	message := dynamic.NewMessage(messageDesc)
	if err := proto.Unmarshal(b, message); err != nil {
		return nil, err
	}
	return message.MarshalJSON()
}
