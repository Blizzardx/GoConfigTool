package common

type ConfigDecoder interface {
	// 将数据转换为字节数组
	Encode(msgObj interface{}) (data interface{}, err error)

	// 将字节数组转换为数据
	Decode(data interface{}, msgObj interface{}) error

	// 编码器的名字
	Name() string

	MimeType() string
}