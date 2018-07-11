package decoder

import "github.com/vmihailenco/msgpack"

type MsgPackDecodeC struct {
}

func (self *MsgPackDecodeC) Name() string {
	return "msgpack"
}
func (self *MsgPackDecodeC) MimeType() string {
	return "application/message-pack"
}

func (self *MsgPackDecodeC) Encode(msgObj interface{}) (data interface{}, err error) {
	return msgpack.Marshal(msgObj)
}

func (self *MsgPackDecodeC) Decode(data interface{}, msgObj interface{}) error {
	return msgpack.Unmarshal(data.([]byte), msgObj)
}
