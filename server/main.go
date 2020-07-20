package main

import (
	"context"
	"log"
	"logRPC/pb"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	port = ":80"
)

type server struct {
	pb.UnimplementedLoggerServer
}

func (s *server) CreateEvent(ctx context.Context, event *pb.Event) (*pb.EventResponse, error) {
	method := event.GetMethod()
	timestamp := event.GetTimestamp()
	log.Printf("Event Logged: %s, at timestamp: %s", method, timestamp)
	eventtimestamp := time.Now().String()
	return &pb.EventResponse{EventLogged: true, Timestamp: eventtimestamp}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoggerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
