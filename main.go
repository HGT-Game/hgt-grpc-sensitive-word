package main

import (
	"fmt"

	"github.com/importcjj/sensitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sensitive_word/protobuf/grpc/wordFilter"
	"sensitive_word/server"
)

var WordFilter = sensitive.New()

func main() {
	WordFilter.LoadWordDict("./txt/wordFilter.txt")
	lis, err := net.Listen("tcp", ":8897")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ss := &server.WordFilterGRPCStruct{S: WordFilter}
	s := grpc.NewServer()
	wordFilter.RegisterWordFilterServer(s, ss)

	reflection.Register(s)
	fmt.Println("注册敏感词过滤grpc服务")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
