package main

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/nats-io/go-nats-streaming"
	"log"
)

func main() {
	packetCh := make(chan []byte, 1)
	packet, _ := utils.GetRandomPacket()
	data, _ := json.Marshal(packet)

	packetCh <- data

	sc, err := stan.Connect("benchmark", "benchmark", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}

	subj, msg := "benchmark", data

	err = sc.Publish(subj, msg)
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", subj, msg)
}
