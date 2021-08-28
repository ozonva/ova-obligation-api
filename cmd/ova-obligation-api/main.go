package main

import (
	"fmt"
	"net"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-obligation-api/internal/repo"
	"github.com/ozonva/ova-obligation-api/internal/server"
	api "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	port       = ":8080"
	configPath = "../../config/.env"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.Get("POSTGRES_HOST"), viper.GetInt("POSTGRES_PORT"), viper.Get("POSTGRES_USER"), viper.Get("POSTGRES_PASSWORD"), viper.Get("POSTGRES_DB"))
	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	if db.Ping() != nil {
		log.Fatal().Err(err)
	}

	defer db.Close()

	log.Info().Msgf("start serve port %s", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	repo := repo.NewObligationRepository(db)

	s := grpc.NewServer()
	api.RegisterObligationRpcServer(s, server.NewObligationRpcServer(&log, repo))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
