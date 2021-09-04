package ova_obligation_producer

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
)

type Event string
type topic string

const (
	Created Event = "created"
	Updated Event = "updated"
	Deleted Event = "deleted"

	createdTopic topic = "ova_obligation_created"
	updatedTopic topic = "ova_obligation_updated"
	deletedTopic topic = "ova_obligation_deleted"
)

var eventToTopicMap map[Event]topic = map[Event]topic{
	Created: createdTopic,
	Updated: updatedTopic,
	Deleted: deletedTopic,
}

type Producer interface {
	Publish(id uint, event Event)
}

type ProducerConfig struct {
	Host string
	Port int
}

type OvaObligationProducer struct {
	done     <-chan bool
	logger   *zerolog.Logger
	config   ProducerConfig
	messages chan *sarama.ProducerMessage
	producer sarama.AsyncProducer
}

func NewProducer(done <-chan bool, logger *zerolog.Logger, config ProducerConfig) (Producer, error) {
	producer := OvaObligationProducer{
		done:     done,
		logger:   logger,
		config:   config,
		messages: make(chan *sarama.ProducerMessage, 5),
	}

	err := producer.init()
	if err != nil {
		return nil, err
	}

	return &producer, nil
}

func (p *OvaObligationProducer) Publish(id uint, event Event) {
	topic := eventToTopicMap[event]
	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic: string(topic),
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(fmt.Sprint(id)),
	}

	p.messages <- msg
}

func (p *OvaObligationProducer) init() error {
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
		defer close(p.messages)
		for {
			select {
			case msg := <-p.messages:
				p.producer.Input() <- msg
				p.logger.Log().Msgf("Produce message: %v", msg)

			case err := <-p.producer.Errors():
				p.logger.Err(err).Msgf("Failed to produce message: %s", err.Error())
			case <-p.done:
				return
			}
		}
	}()

	return nil
}
