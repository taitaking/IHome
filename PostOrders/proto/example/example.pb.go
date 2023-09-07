// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostOrders

import (
	fmt "fmt"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Sessionid            string   `protobuf:"bytes,1,opt,name=Sessionid,proto3" json:"Sessionid,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Errno                string   `protobuf:"bytes,2,opt,name=Errno,proto3" json:"Errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,3,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
	OrderId              int64    `protobuf:"varint,1,opt,name=Order_id,json=OrderId,proto3" json:"Order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostOrders.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostOrders.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbd, 0x4b, 0xc4, 0x40,
	0x10, 0xc5, 0x8d, 0xd1, 0x7c, 0x0c, 0x56, 0x83, 0x68, 0xfc, 0x28, 0x62, 0xaa, 0x54, 0x2b, 0x68,
	0x69, 0x27, 0xa4, 0xb0, 0x52, 0x36, 0xbd, 0xa2, 0xee, 0x10, 0x02, 0x26, 0x13, 0x67, 0x72, 0xc7,
	0xdd, 0x7f, 0x7f, 0xb0, 0xc9, 0x91, 0xea, 0xae, 0x9a, 0x79, 0x8f, 0x07, 0xf3, 0x7b, 0x03, 0x77,
	0x83, 0xf0, 0xc8, 0x8f, 0xb4, 0xf9, 0xee, 0x86, 0x3f, 0xda, 0x4f, 0xe3, 0x5d, 0xbc, 0x6e, 0xd8,
	0x74, 0xed, 0xaf, 0xb0, 0x51, 0x59, 0x9b, 0x0f, 0xd6, 0xf1, 0x5d, 0x1c, 0x89, 0x16, 0x2f, 0x10,
	0x5b, 0xfa, 0x5f, 0x91, 0x8e, 0x78, 0x0f, 0x69, 0x4d, 0xaa, 0x2d, 0xf7, 0xad, 0xcb, 0x82, 0x3c,
	0x28, 0x53, 0xbb, 0x18, 0x88, 0x70, 0xf6, 0xca, 0x6e, 0x9b, 0x9d, 0xe6, 0x41, 0x79, 0x61, 0xfd,
	0x5e, 0xd4, 0x90, 0x58, 0xd2, 0x81, 0x7b, 0x25, 0xbc, 0x84, 0xf3, 0x4a, 0xa4, 0x67, 0x1f, 0x48,
	0xed, 0x24, 0xf0, 0x0a, 0xa2, 0x4a, 0xa4, 0xd3, 0x26, 0x0b, 0xbd, 0x3d, 0x2b, 0xbc, 0x81, 0xc4,
	0x03, 0x7c, 0xcd, 0xa7, 0x42, 0x1b, 0x7b, 0xfd, 0xe6, 0x9e, 0x3e, 0x21, 0xae, 0x26, 0x76, 0xac,
	0x01, 0x16, 0x54, 0xcc, 0xcd, 0x81, 0x12, 0x66, 0x6e, 0x70, 0xfb, 0x70, 0x24, 0x31, 0x61, 0x16,
	0x27, 0x3f, 0x91, 0xff, 0xc8, 0xf3, 0x2e, 0x00, 0x00, 0xff, 0xff, 0x95, 0x9a, 0x6c, 0x7e, 0x30,
	0x01, 0x00, 0x00,
}
