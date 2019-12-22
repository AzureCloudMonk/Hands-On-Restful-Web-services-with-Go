// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/encryption.proto

package encryption

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Encrypter service

type EncrypterService interface {
	Encrypt(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Decrypt(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type encrypterService struct {
	c    client.Client
	name string
}

func NewEncrypterService(name string, c client.Client) EncrypterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "encrypter"
	}
	return &encrypterService{
		c:    c,
		name: name,
	}
}

func (c *encrypterService) Encrypt(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Encrypter.Encrypt", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encrypterService) Decrypt(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Encrypter.Decrypt", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Encrypter service

type EncrypterHandler interface {
	Encrypt(context.Context, *Request, *Response) error
	Decrypt(context.Context, *Request, *Response) error
}

func RegisterEncrypterHandler(s server.Server, hdlr EncrypterHandler, opts ...server.HandlerOption) error {
	type encrypter interface {
		Encrypt(ctx context.Context, in *Request, out *Response) error
		Decrypt(ctx context.Context, in *Request, out *Response) error
	}
	type Encrypter struct {
		encrypter
	}
	h := &encrypterHandler{hdlr}
	return s.Handle(s.NewHandler(&Encrypter{h}, opts...))
}

type encrypterHandler struct {
	EncrypterHandler
}

func (h *encrypterHandler) Encrypt(ctx context.Context, in *Request, out *Response) error {
	return h.EncrypterHandler.Encrypt(ctx, in, out)
}

func (h *encrypterHandler) Decrypt(ctx context.Context, in *Request, out *Response) error {
	return h.EncrypterHandler.Decrypt(ctx, in, out)
}