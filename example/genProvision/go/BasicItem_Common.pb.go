// Code generated by protoc-gen-go. DO NOT EDIT.
// source: BasicItem_Common.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	BasicItem_Common.proto

It has these top-level messages:
	BasicItem_CommonConfig
	BasicItem_CommonLineInfo
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BasicItem_CommonConfig struct {
	Content []*BasicItem_CommonLineInfo `protobuf:"bytes,1,rep,name=Content" json:"Content,omitempty"`
}

func (m *BasicItem_CommonConfig) Reset()                    { *m = BasicItem_CommonConfig{} }
func (m *BasicItem_CommonConfig) String() string            { return proto.CompactTextString(m) }
func (*BasicItem_CommonConfig) ProtoMessage()               {}
func (*BasicItem_CommonConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BasicItem_CommonConfig) GetContent() []*BasicItem_CommonLineInfo {
	if m != nil {
		return m.Content
	}
	return nil
}

type BasicItem_CommonLineInfo struct {
	Id          int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Icon        string `protobuf:"bytes,2,opt,name=Icon" json:"Icon,omitempty"`
	Quality     int32  `protobuf:"varint,3,opt,name=Quality" json:"Quality,omitempty"`
	Price       int32  `protobuf:"varint,4,opt,name=Price" json:"Price,omitempty"`
	LimitNum    int32  `protobuf:"varint,5,opt,name=LimitNum" json:"LimitNum,omitempty"`
	Acauire     int32  `protobuf:"varint,6,opt,name=Acauire" json:"Acauire,omitempty"`
	ConsumeItem string `protobuf:"bytes,7,opt,name=ConsumeItem" json:"ConsumeItem,omitempty"`
	ConsumeCoin int32  `protobuf:"varint,8,opt,name=ConsumeCoin" json:"ConsumeCoin,omitempty"`
	FormatIndex int32  `protobuf:"varint,9,opt,name=FormatIndex" json:"FormatIndex,omitempty"`
}

func (m *BasicItem_CommonLineInfo) Reset()                    { *m = BasicItem_CommonLineInfo{} }
func (m *BasicItem_CommonLineInfo) String() string            { return proto.CompactTextString(m) }
func (*BasicItem_CommonLineInfo) ProtoMessage()               {}
func (*BasicItem_CommonLineInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BasicItem_CommonLineInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *BasicItem_CommonLineInfo) GetQuality() int32 {
	if m != nil {
		return m.Quality
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetLimitNum() int32 {
	if m != nil {
		return m.LimitNum
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetAcauire() int32 {
	if m != nil {
		return m.Acauire
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetConsumeItem() string {
	if m != nil {
		return m.ConsumeItem
	}
	return ""
}

func (m *BasicItem_CommonLineInfo) GetConsumeCoin() int32 {
	if m != nil {
		return m.ConsumeCoin
	}
	return 0
}

func (m *BasicItem_CommonLineInfo) GetFormatIndex() int32 {
	if m != nil {
		return m.FormatIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*BasicItem_CommonConfig)(nil), "config.BasicItem_CommonConfig")
	proto.RegisterType((*BasicItem_CommonLineInfo)(nil), "config.BasicItem_CommonLineInfo")
}

func init() { proto.RegisterFile("BasicItem_Common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xda, 0x24, 0xed, 0x16, 0x3c, 0x0c, 0x22, 0x83, 0xa7, 0xd0, 0x53, 0x4e, 0x39,
	0xe8, 0xcd, 0x9b, 0x2e, 0x08, 0x0b, 0x45, 0x34, 0x78, 0x97, 0xb8, 0xd9, 0xca, 0x80, 0x3b, 0x23,
	0xe9, 0x06, 0xf4, 0x15, 0x7c, 0x6a, 0xc9, 0xc6, 0x4a, 0x50, 0xbc, 0xed, 0xff, 0xcf, 0xf7, 0x0d,
	0xc3, 0xaa, 0xb3, 0x9b, 0xf6, 0x40, 0xd6, 0x04, 0xe7, 0x9f, 0xb4, 0x78, 0x2f, 0x5c, 0xbf, 0xf5,
	0x12, 0x04, 0x72, 0x2b, 0xbc, 0xa7, 0x97, 0xed, 0xe3, 0x5f, 0x42, 0xc7, 0x09, 0x5c, 0xa9, 0x42,
	0x0b, 0x07, 0xc7, 0x01, 0x93, 0x72, 0x51, 0x6d, 0x2e, 0xca, 0x7a, 0x72, 0xea, 0xdf, 0xc2, 0x8e,
	0xd8, 0x19, 0xde, 0x4b, 0x73, 0x14, 0xb6, 0x9f, 0xa9, 0xc2, 0xff, 0x28, 0x38, 0x51, 0xa9, 0xe9,
	0x30, 0x29, 0x93, 0x2a, 0x6b, 0x52, 0xd3, 0x01, 0xa8, 0xa5, 0xb1, 0xc2, 0x98, 0x96, 0x49, 0xb5,
	0x6e, 0xe2, 0x1b, 0x50, 0x15, 0x0f, 0x43, 0xfb, 0x4a, 0xe1, 0x03, 0x17, 0x11, 0x3c, 0x46, 0x38,
	0x55, 0xd9, 0x7d, 0x4f, 0xd6, 0xe1, 0x32, 0xf6, 0x53, 0x80, 0x73, 0xb5, 0xda, 0x91, 0xa7, 0x70,
	0x37, 0x78, 0xcc, 0xe2, 0xe0, 0x27, 0x8f, 0xbb, 0xae, 0x6d, 0x3b, 0x50, 0xef, 0x30, 0x9f, 0x76,
	0x7d, 0x47, 0x28, 0xd5, 0x46, 0x0b, 0x1f, 0x06, 0xef, 0xc6, 0x3b, 0xb1, 0x88, 0x07, 0xcc, 0xab,
	0x19, 0xa1, 0x85, 0x18, 0x57, 0xd1, 0x9f, 0x57, 0x23, 0x71, 0x2b, 0xbd, 0x6f, 0x83, 0xe1, 0xce,
	0xbd, 0xe3, 0x7a, 0x22, 0x66, 0xd5, 0x73, 0x1e, 0x7f, 0xfc, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x46, 0x49, 0x6e, 0x8b, 0x01, 0x00, 0x00,
}
