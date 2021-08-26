package main

import (
	"net"
	"os"
	"time"

	"github.com/ozonva/ova-obligation-api/internal/server"
	api "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Msgf("start serve port %s", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterObligationRpcServer(s, server.NewObligationRpcServer(&log))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
