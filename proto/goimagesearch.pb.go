// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/goimagesearch.proto

package goimagesearch

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

type PingRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_goimagesearch_9159739cb308a27c, []int{0}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PingResponse struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_goimagesearch_9159739cb308a27c, []int{1}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
}

func init() {
	proto.RegisterFile("proto/goimagesearch.proto", fileDescriptor_goimagesearch_9159739cb308a27c)
}

var fileDescriptor_goimagesearch_9159739cb308a27c = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcf, 0xcf, 0xcc, 0x4d, 0x4c, 0x4f, 0x2d, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0xd0,
	0x03, 0x8b, 0x29, 0x29, 0x72, 0x71, 0x07, 0x64, 0xe6, 0xa5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x81, 0xd9, 0x4a, 0x1a, 0x5c, 0x3c, 0x10, 0x25, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x12,
	0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x65, 0x30, 0xae, 0x91,
	0x19, 0x17, 0xaf, 0x7b, 0xbe, 0x27, 0xc8, 0x8e, 0x60, 0xb0, 0x1d, 0x42, 0xaa, 0x5c, 0x2c, 0x20,
	0xad, 0x42, 0x3c, 0x7a, 0x48, 0x96, 0x48, 0xf1, 0xea, 0x21, 0x9b, 0xa7, 0xc4, 0x90, 0xc4, 0x06,
	0x76, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xbe, 0xb3, 0x70, 0xa8, 0x00, 0x00, 0x00,
}
