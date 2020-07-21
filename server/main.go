package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
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

func (s *server) CreateEvent(ctx context.Context, event *pb.Event) (*pb.EventResponse, error) {
	method := event.GetMethod()
	timestamp := event.GetTimestamp()
	SaveToFile(ctx, event)
	log.Printf("Event Logged: %s, at timestamp: %s", method, timestamp)
	eventtimestamp := time.Now().String()
	return &pb.EventResponse{EventLogged: true, Timestamp: eventtimestamp}, nil
}

// SaveToFile saves to a json file
func SaveToFile(ctx context.Context, e *pb.Event) (bool, error) {
	f, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
		return false, err
	}
	err = ioutil.WriteFile("logtest.json", f, 0644)
	if err != nil {
		log.Fatalf("failed to write to file: %v", err)
		return false, err
	}
	return true, nil
}
