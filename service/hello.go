package service

import (
	"article-db/pb"
	"context"
	"trpc.group/trpc-go/trpc-go/log"
)

func (ServiceImpl) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Infof("got hello request: %s", req.Msg)
	return &pb.HelloResponse{Msg: "Hello " + req.Msg + "!"}, nil
}
