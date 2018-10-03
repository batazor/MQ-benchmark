package amqp

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/pkg/amqp"
	"github.com/batazor/MQ-benchmark/utils"
	"testing"
)

var (
	packet, _ = utils.GetRandomPacket()
	data, _   = json.Marshal(packet)
)

func BenchmarkReceiver(b *testing.B) {
	CONSUMER := amqp.Listen()

	for i := 0; i < 100000; i++ {
		CONSUMER.Publish(data)
	}
}

//func BenchmarkConsumer(b *testing.B) {
//	CONSUMER := amqp.Listen()
//	deliveries, _ := CONSUMER.AnnounceQueue()
//
//	go CONSUMER.Handle(deliveries, handler, CONSUMER.QueueName)
//}
//
//func handler(deliveries <-chan AMQP.Delivery) {
//	for d := range deliveries {
//		go d.Ack(false)
//	}
//}
