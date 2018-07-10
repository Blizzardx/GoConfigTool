package decoder

import (
	"github.com/gogo/protobuf/proto"
)

type PBDecodeC struct {
}

// 编码器的名称
func (self *PBDecodeC) Name() string {
	return "gogopb"
}

func (self *PBDecodeC) MimeType() string {
	return "application/x-protobuf"
}

func (self *PBDecodeC) Encode(msgObj interface{}) (data interface{}, err error) {

	return proto.Marshal(msgObj.(proto.Message))

}

func (self *PBDecodeC) Decode(data interface{}, msgObj interface{}) error {

	return proto.Unmarshal(data.([]byte), msgObj.(proto.Message))
}