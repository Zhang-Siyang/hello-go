package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello/pkg/zaps"
	pb "hello/protos"
	"net"
)

type server struct {
	pb.UnimplementedJackieChanServer
}

func (*server) GetLonger(c context.Context, req *pb.LongerRequest) (resp *pb.LongerReply, err error) {
	zaps.Logger(c).Info(fmt.Sprintf("get req %+v", req))
	if len(req.GetM()) >= len(req.GetN()) {
		return &pb.LongerReply{Longer: req.GetM()}, nil
	}
	return &pb.LongerReply{Longer: req.GetN()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3001")
	if err != nil {
		zaps.GlobalLogger().Fatal(fmt.Sprintf("failed to listen: %v", err))
	}
	s := grpc.NewServer()
	pb.RegisterJackieChanServer(s, &server{})
	zaps.GlobalLogger().Info(fmt.Sprintf("server listening at %v", lis.Addr()))
	if err := s.Serve(lis); err != nil {
		zaps.GlobalLogger().Fatal(fmt.Sprintf("failed to serve: %v", err))
	}
}
