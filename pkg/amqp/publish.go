package amqp

import (
	"github.com/streadway/amqp"
	"time"
)

func (c *Consumer) Publish(message []byte) error {

	err := c.channel.Publish(
		c.QueueName,
		c.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Transient,
			Body:         message,
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		log.Info("Failed to bind a queue: ", err)
	}

	return nil
}
