package main

import (
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strings"
)

var (
	log          = logrus.New()
	exchangeList = strings.Split(AMQP_EXCHANGE_LIST, ",")

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://telemetry:telemetry@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "go-logger-packets")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "demo1, demo2")
	AMQP_EXCHANGE_TYPE = utils.Getenv("AMQP_EXCHANGE_TYPE", "headers")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue

	done     chan error
	packetCh chan []byte

	consumerTag  string // Name that consumer identifies itself to the server with
	uri          string // uri of the rabbitmq server
	changes      string // exchange that we will bind to
	exchangeType string // topic, direct, etc...
	bindingKey   string // routing key that we are using
}

func NewConsumer(uri, changes, exchangeType, queueName, bindingKey, consumerTag string) *Consumer {
	return &Consumer{
		uri:          uri,
		changes:      changes,
		bindingKey:   bindingKey,
		exchangeType: exchangeType,
		conn:         nil,
		channel:      nil,
		consumerTag:  consumerTag,
		done:         make(chan error),
	}

}

func (s *Consumer) Connect() {
	var err error
	s.conn, err = amqp.Dial(AMQP_API)
	if err != nil {
		log.Info("Failed to connect to RabbitMQ: ", err)
	}

	s.channel, err = s.conn.Channel()
	if err != nil {
		log.Info("Failed to open a channel: ", err)
	}

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")
		err = s.channel.ExchangeDeclare(
			name,
			AMQP_EXCHANGE_TYPE,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Info("Failed to declare the Exchange: ", err)
		}
	}

	s.queue, err = s.channel.QueueDeclare(
		AMQP_NAME_QUEUE,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Info("Failed to declare a queue: ", err)
	}

	for _, echangeName := range exchangeList {
		name := strings.Trim(echangeName, " ")

		err := s.channel.QueueBind(
			s.queue.Name,
			"",
			name,
			false,
			nil,
		)
		if err != nil {
			log.Info("Failed to bind a queue: ", err)
		}
	}
}
