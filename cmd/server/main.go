package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	greetv1 "github.com/abitofhelp/bazel-go-googleapis/gen/greet/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = ":8080"
)

// server is used to implement helloworld.GreeterServer.
type GreetServer struct {
	greetv1.UnsafeGreetServiceServer
}

func (x *GreetServer) Greet(ctx context.Context, req *greetv1.GreetRequest) (*greetv1.GreetResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal().Msgf("failed to convert the request to a JSON: %#v", err)
	}

	log.Info().Msgf("Received request: '%s'", data)
	res := &greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Name),
		Sent:     timestamppb.New(time.Now()),
	}
	return res, nil
}

func main() {

	greeter := &GreetServer{}
	s := grpc.NewServer()
	greetv1.RegisterGreetServiceServer(s, greeter)
	// Register reflection service on gRPC server.
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %#v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
