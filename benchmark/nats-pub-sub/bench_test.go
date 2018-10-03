package main

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/nats-io/go-nats"
	"log"
	"testing"
)

var (
	packet, _ = utils.GetRandomPacket()
	data, _   = json.Marshal(packet)
)

func BenchmarkReceiver(b *testing.B) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 100000; i++ {
		nc.Publish("benchmark", data)
	}
}
