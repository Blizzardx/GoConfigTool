package decoder

import "encoding/json"

type JsonDecodeC struct {
}

// 编码器的名称
func (self *JsonDecodeC) Name() string {
	return "json"
}

func (self *JsonDecodeC) MimeType() string {
	return "application/json"
}

// 将结构体编码为JSON的字节数组
func (self *JsonDecodeC) Encode(msgObj interface{}) (data interface{}, err error) {

	return json.Marshal(msgObj)

}

// 将JSON的字节数组解码为结构体
func (self *JsonDecodeC) Decode(data interface{}, msgObj interface{}) error {

	return json.Unmarshal(data.([]byte), msgObj)
}
