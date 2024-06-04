package main

import (
	// 	"context"
	// 	"log"
	// 	"net/http"
	// 	"time"
	"context"
	"os"
	"time"

	greetv1 "github.com/abitofhelp/bazel-go-googleapis/gen/greet/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	address     = "localhost:8080"
	defaultName = "Jane"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greetv1.NewGreetServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &greetv1.GreetRequest{Name: name})
	if err != nil {
		log.Fatal().Msgf("could not greet: %v", err)
	}
	log.Info().Msg("Greeting: " + r.GetGreeting())
}

func buildRequest(name string) *greetv1.GreetRequest {
	now := time.Now().UTC()
	timestamp := timestamppb.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.Nanosecond()),
	}

	return &greetv1.GreetRequest{
		Name:    name,
		Sent: &timestamp,
	}
}
