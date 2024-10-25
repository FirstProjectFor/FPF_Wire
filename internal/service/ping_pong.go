package service

import (
	v1 "FPF_Wire/api/helloworld/v1"
	context "context"
	"fmt"
)

var _ v1.PingPongHTTPServer = (*PingPong)(nil)

func NewPingPong() *PingPong {
	return &PingPong{}
}

type PingPong struct {
	v1.UnimplementedPingPongServer
}

func (p PingPong) Ping(ctx context.Context, request *v1.PingRequest) (*v1.PingReply, error) {
	fmt.Println(request.String())
	return &v1.PingReply{Message: "pong"}, nil
}
