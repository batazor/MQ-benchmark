package main

import (
	"encoding/json"
	"flag"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	log = logrus.New()
)

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")

	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}
	defer nc.Close()

	packetCh := make(chan []byte, 1)
	packet, _ := utils.GetRandomPacket()
	data, _ := json.Marshal(packet)

	packetCh <- data

	for {
		select {
		case <-packetCh:
			for i := 0; i < 100000; i++ {
				go nc.Request("benchmark", data, 100*time.Millisecond)
			}

			//logrus.Info("json", string(data))

			time.Sleep(time.Second * 1)
			// logrus.Info("Add 1000 packets")

			p, _ := utils.GetRandomPacket()
			data, _ = json.Marshal(p)
			packetCh <- data
		}
	}
}
