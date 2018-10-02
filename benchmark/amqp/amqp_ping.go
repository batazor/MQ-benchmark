package main

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/pkg/amqp"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	log = logrus.New()
)

func main() {
	packetCh := make(chan []byte, 1)
	packet, _ := utils.GetRandomPacket()
	data, _ := json.Marshal(packet)

	CONSUMER := amqp.Listen()

	packetCh <- data

	for {
		select {
		case <-packetCh:
			for i := 0; i < 1000; i++ {
				go CONSUMER.Publish(data)
			}

			//logrus.Info("json", string(data))

			time.Sleep(time.Millisecond * 100)
			// logrus.Info("Add 1000 packets")

			p, _ := utils.GetRandomPacket()
			data, _ = json.Marshal(p)
			packetCh <- data
		}
	}
}
