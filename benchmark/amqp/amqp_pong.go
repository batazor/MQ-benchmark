package main

import (
	"github.com/batazor/MQ-benchmark/pkg/amqp"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/sirupsen/logrus"
	AMQP "github.com/streadway/amqp"
)

var (
	log = logrus.New()
)

func main() {
	CONSUMER := amqp.Listen()

	deliveries, err := CONSUMER.AnnounceQueue()
	if err != nil {
		log.Warn(err)
	}

	CONSUMER.Handle(deliveries, handler, CONSUMER.QueueName)
}

func handler(deliveries <-chan AMQP.Delivery) {
	threads := utils.MaxParallelism()

	for i := 0; i < threads; i++ {
		go func() {
			for d := range deliveries {
				d.Ack(false)
			}
		}()
	}
}
