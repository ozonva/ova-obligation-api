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
	queue "github.com/ozonva/ova-obligation-api/pkg/ova-obligation-producer"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Config struct {
	Port int `mapstructure:"PORT"`

	DbPort     int    `mapstructure:"POSTGRES_PORT"`
	DbHost     string `mapstructure:"POSTGRES_HOST"`
	DbUser     string `mapstructure:"POSTGRES_USER"`
	DbPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DbName     string `mapstructure:"POSTGRES_DB"`

	KafkaHost string `mapstructure:"KAFKA_HOST"`
	KafkaPort int    `mapstructure:"KAFKA_PORT"`
}

const configPath = ".env"

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	if db.Ping() != nil {
		log.Fatal().Err(err)
	}

	defer db.Close()

	log.Info().Msgf("start serve port %d", config.Port)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	done := make(chan bool)
	producer, err := queue.NewProducer(done, &log, queue.ProducerConfig{
		Host: config.KafkaHost,
		Port: config.KafkaPort,
	})

	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	repo := repo.NewObligationRepository(db)

	s := grpc.NewServer()
	api.RegisterObligationRpcServer(s, server.NewObligationRpcServer(&log, repo, producer))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
