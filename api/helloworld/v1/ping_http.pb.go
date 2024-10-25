// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v3.21.8
// source: helloworld/v1/ping.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPingPongPing = "/helloworld.v1.PingPong/Ping"

type PingPongHTTPServer interface {
	// Ping Sends a greeting
	Ping(context.Context, *EmptyRequest) (*PingReply, error)
}

func RegisterPingPongHTTPServer(s *http.Server, srv PingPongHTTPServer) {
	r := s.Route("/")
	r.GET("/ping", _PingPong_Ping0_HTTP_Handler(srv))
}

func _PingPong_Ping0_HTTP_Handler(srv PingPongHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EmptyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPingPongPing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*EmptyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingReply)
		return ctx.Result(200, reply)
	}
}

type PingPongHTTPClient interface {
	Ping(ctx context.Context, req *EmptyRequest, opts ...http.CallOption) (rsp *PingReply, err error)
}

type PingPongHTTPClientImpl struct {
	cc *http.Client
}

func NewPingPongHTTPClient(client *http.Client) PingPongHTTPClient {
	return &PingPongHTTPClientImpl{client}
}

func (c *PingPongHTTPClientImpl) Ping(ctx context.Context, in *EmptyRequest, opts ...http.CallOption) (*PingReply, error) {
	var out PingReply
	pattern := "/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPingPongPing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}