package main

import (
	"article-db/pb"
	"article-db/service"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterGreeterService(s, &service.Greeter{})
	if err := s.Serve(); err != nil {
		log.Error(err)
	}
}
