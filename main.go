package main

import (
	"fmt"
	pb "github.com/anthony-bible/grpc-example/proto/example"
		"google.golang.org/grpc"
	"net"
	"log"
	"context"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/metadata"
)

type YourService struct{
  pb.UnimplementedYourServiceServer
}

func (s *YourService) YourMethod(ctx context.Context, e *emptypb.Empty) (*pb.YourResponse, error) {
	// print grpc metadata like headers
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	return &pb.YourResponse{Result: "Hello "  }, nil
}


func main() {
	fmt.Println("Hello, World!")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:50051", ))
if err != nil {
  log.Fatalf("failed to listen: %v", err)
}
s := grpc.NewServer()
pb.RegisterYourServiceServer(s , &YourService{})
reflection.Register(s)
s.Serve(lis)
}
