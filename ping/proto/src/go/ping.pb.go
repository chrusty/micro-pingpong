// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

/*
Package com_cruft_service_ping is a generated protocol buffer package.

It is generated from these files:
	ping.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package com_cruft_service_ping

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type PingRequest struct {
	Iteration int32 `protobuf:"varint,1,opt,name=iteration" json:"iteration,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

type PingResponse struct {
	Noise     string `protobuf:"bytes,1,opt,name=noise" json:"noise,omitempty"`
	Iteration int32  `protobuf:"varint,2,opt,name=iteration" json:"iteration,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetNoise() string {
	if m != nil {
		return m.Noise
	}
	return ""
}

func (m *PingResponse) GetIteration() int32 {
	if m != nil {
		return m.Iteration
	}
	return 0
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "com.cruft.service.ping.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "com.cruft.service.ping.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Ping service

type PingClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PingResponse, error)
}

type pingClient struct {
	c           client.Client
	serviceName string
}

func NewPingClient(serviceName string, c client.Client) PingClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.cruft.service.ping"
	}
	return &pingClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *pingClient) Ping(ctx context.Context, in *PingRequest, opts ...client.CallOption) (*PingResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Ping.Ping", in)
	out := new(PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingHandler interface {
	Ping(context.Context, *PingRequest, *PingResponse) error
}

func RegisterPingHandler(s server.Server, hdlr PingHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Ping{hdlr}, opts...))
}

type Ping struct {
	PingHandler
}

func (h *Ping) Ping(ctx context.Context, in *PingRequest, out *PingResponse) error {
	return h.PingHandler.Ping(ctx, in, out)
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xc8, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xce, 0xcf, 0xd5, 0x4b, 0x2e, 0x2a, 0x4d,
	0x2b, 0xd1, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0xc9, 0x2a, 0x69, 0x73, 0x71,
	0x07, 0x64, 0xe6, 0xa5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xc9, 0x70, 0x71, 0x66,
	0x96, 0xa4, 0x16, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x21,
	0x04, 0x94, 0x9c, 0xb8, 0x78, 0x20, 0x8a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8,
	0x58, 0xf3, 0xf2, 0x33, 0x8b, 0x53, 0xc1, 0x2a, 0x39, 0x83, 0x20, 0x1c, 0x54, 0x33, 0x98, 0xd0,
	0xcc, 0x30, 0x8a, 0xe6, 0x62, 0x01, 0x99, 0x21, 0x14, 0x0c, 0xa5, 0x95, 0xf5, 0xb0, 0xbb, 0x4c,
	0x0f, 0xc9, 0x59, 0x52, 0x2a, 0xf8, 0x15, 0x41, 0x9c, 0xa3, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0xac,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x8b, 0x1c, 0x3c, 0xfa, 0x00, 0x00, 0x00,
}