package ova_obligation_producer

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
)

type Event string

const (
	Created Event = "created"
	Updated Event = "updated"
	Deleted Event = "deleted"

	topic = "ova_obligation"
)

type Producer interface {
	Publish(id uint, event Event)
}

type ProducerConfig struct {
	Host string
	Port int
}

type OvaObligationProducer struct {
	logger   *zerolog.Logger
	config   ProducerConfig
	producer sarama.AsyncProducer
}

func NewProducer(ctx context.Context, logger *zerolog.Logger, config ProducerConfig) (Producer, error) {
	producer := OvaObligationProducer{
		logger: logger,
		config: config,
	}

	err := producer.init(ctx)
	if err != nil {
		return nil, err
	}

	return &producer, nil
}

func (p *OvaObligationProducer) Publish(id uint, event Event) {
	msg := &sarama.ProducerMessage{
		Topic: string(topic),
		Key:   sarama.StringEncoder(event),
		Value: sarama.StringEncoder(fmt.Sprint(id)),
	}

	p.producer.Input() <- msg
	p.logger.Log().Msgf("Produce message: key: %s, value: %s", event, fmt.Sprint(id))
}

func (p *OvaObligationProducer) init(ctx context.Context) error {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll

	brokers := []string{fmt.Sprintf("%s:%d", p.config.Host, p.config.Port)}

	var err error
	p.producer, err = sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return err
	}

	go func() {
		defer p.producer.Close()
		for {
			select {
			case err := <-p.producer.Errors():
				p.logger.Err(err).Msgf("Failed to produce message: %s", err.Error())
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
