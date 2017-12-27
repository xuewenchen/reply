package service

import (
	"context"
	pb "kit/model/example"
)

type helloServer struct {
	svr *service
}

func (h *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}
