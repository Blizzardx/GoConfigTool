package decoder

import "encoding/xml"

type XmlDecodeC struct {
}

// 编码器的名称
func (self *XmlDecodeC) Name() string {
	return "xml"
}

func (self *XmlDecodeC) MimeType() string {
	return "application/xml"
}

func (self *XmlDecodeC) Encode(msgObj interface{}) (data interface{}, err error) {

	return xml.Marshal(msgObj)

}

func (self *XmlDecodeC) Decode(data interface{}, msgObj interface{}) error {

	return xml.Unmarshal(data.([]byte), msgObj)
}
