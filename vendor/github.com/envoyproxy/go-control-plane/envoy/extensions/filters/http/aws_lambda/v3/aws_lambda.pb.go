// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto

package envoy_extensions_filters_http_aws_lambda_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Config struct {
	Arn                  string   `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
	PayloadPassthrough   bool     `protobuf:"varint,2,opt,name=payload_passthrough,json=payloadPassthrough,proto3" json:"payload_passthrough,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetArn() string {
	if m != nil {
		return m.Arn
	}
	return ""
}

func (m *Config) GetPayloadPassthrough() bool {
	if m != nil {
		return m.PayloadPassthrough
	}
	return false
}

type PerRouteConfig struct {
	InvokeConfig         *Config  `protobuf:"bytes,1,opt,name=invoke_config,json=invokeConfig,proto3" json:"invoke_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a22f80ea45c04127, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

func (m *PerRouteConfig) GetInvokeConfig() *Config {
	if m != nil {
		return m.InvokeConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "envoy.extensions.filters.http.aws_lambda.v3.Config")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.extensions.filters.http.aws_lambda.v3.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/aws_lambda/v3/aws_lambda.proto", fileDescriptor_a22f80ea45c04127)
}

var fileDescriptor_a22f80ea45c04127 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x49, 0x94, 0xaa, 0xa7, 0x96, 0x12, 0x07, 0x6b, 0x41, 0xa9, 0x9d, 0x0a, 0xc2, 0x1d,
	0x34, 0x2e, 0x2d, 0x5d, 0x1a, 0x57, 0x87, 0x10, 0x10, 0xdc, 0xca, 0xab, 0xb9, 0x36, 0x87, 0xf1,
	0xee, 0xb8, 0xbb, 0xa4, 0xed, 0xe6, 0xe8, 0xe6, 0xee, 0x77, 0xf0, 0x0b, 0xb8, 0x0b, 0xae, 0x7e,
	0x1d, 0x27, 0x49, 0x2e, 0xd2, 0x4a, 0xa7, 0x6e, 0x79, 0xf9, 0xbf, 0xff, 0xff, 0xfd, 0xde, 0x3d,
	0x34, 0xa4, 0x3c, 0x17, 0x4b, 0x42, 0x17, 0x86, 0x72, 0xcd, 0x04, 0xd7, 0x64, 0xca, 0x52, 0x43,
	0x95, 0x26, 0x89, 0x31, 0x92, 0xc0, 0x5c, 0x8f, 0x53, 0x78, 0x9a, 0xc4, 0x40, 0x72, 0x7f, 0xad,
	0xc2, 0x52, 0x09, 0x23, 0xbc, 0xab, 0xd2, 0x8d, 0x57, 0x6e, 0x5c, 0xb9, 0x71, 0xe1, 0xc6, 0x6b,
	0xfd, 0xb9, 0xdf, 0x3a, 0xcf, 0x62, 0x09, 0x04, 0x38, 0x17, 0x06, 0x4c, 0x39, 0x4a, 0x1b, 0x30,
	0x99, 0xb6, 0x59, 0xad, 0xcb, 0x0d, 0x39, 0xa7, 0xaa, 0x08, 0x65, 0x7c, 0x56, 0xb5, 0x9c, 0xe6,
	0x90, 0xb2, 0x18, 0x0c, 0x25, 0x7f, 0x1f, 0x56, 0xe8, 0xbc, 0x3a, 0xa8, 0x76, 0x23, 0xf8, 0x94,
	0xcd, 0xbc, 0x33, 0xb4, 0x03, 0x8a, 0x37, 0x9d, 0xb6, 0xd3, 0x3d, 0x08, 0xf6, 0x7e, 0x82, 0x5d,
	0xe5, 0x36, 0x9c, 0xa8, 0xf8, 0xe7, 0x11, 0x74, 0x22, 0x61, 0x99, 0x0a, 0x88, 0xc7, 0x12, 0xb4,
	0x36, 0x89, 0x12, 0xd9, 0x2c, 0x69, 0xba, 0x6d, 0xa7, 0xbb, 0x1f, 0x79, 0x95, 0x14, 0xae, 0x94,
	0x41, 0xff, 0xed, 0xf3, 0xe5, 0xe2, 0x1a, 0xf5, 0xec, 0x96, 0x0f, 0xe5, 0x80, 0x6a, 0xc3, 0xcd,
	0x05, 0x7b, 0x90, 0xca, 0x04, 0xb0, 0xc5, 0xe8, 0xbc, 0x3b, 0xa8, 0x1e, 0x52, 0x15, 0x89, 0xcc,
	0xd0, 0x8a, 0xec, 0x1e, 0x1d, 0x33, 0x9e, 0x8b, 0x47, 0x3a, 0xb6, 0x49, 0x25, 0xe3, 0x61, 0xcf,
	0xc7, 0x5b, 0x3c, 0x62, 0x15, 0x1f, 0x1d, 0xd9, 0x24, 0x5b, 0x0d, 0x46, 0x05, 0xe7, 0x10, 0x0d,
	0xb6, 0xe1, 0xfc, 0x0f, 0x17, 0xdc, 0x7d, 0x3c, 0x7f, 0x7d, 0xd7, 0xdc, 0x86, 0x8b, 0xfa, 0x4c,
	0x58, 0x22, 0xa9, 0xc4, 0x62, 0xb9, 0x0d, 0x5c, 0x50, 0x1f, 0xcd, 0xf5, 0x6d, 0x59, 0x85, 0xc5,
	0x59, 0x42, 0x67, 0x52, 0x2b, 0xef, 0xe3, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x1f, 0x81,
	0x59, 0x67, 0x02, 0x00, 0x00,
}
