package main

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/nsqio/go-nsq"
	"testing"
)

var (
	packet, _ = utils.GetRandomPacket()
	data, _   = json.Marshal(packet)
)

func BenchmarkReceiver(b *testing.B) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 100000; i++ {
		w.Publish("write_test", data)
	}
}
