package service

import (
	"context"
	opentracing "github.com/opentracing/opentracing-go"
	pb "kit/model/example"
	"reply/model"
)

type helloServer struct {
	svr *service
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	span := opentracing.SpanFromContext(ctx)
	span.LogFields()

	reply := &model.Reply{
		SourceId: int64(1),
		TypeId:   int8(0),
		Mid:      int64(123),
		Comment:  "这是一段文字",
		ParentId: int64(0),
	}
	s.svr.Add(ctx, reply)
	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}
