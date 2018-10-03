package main

import (
	"encoding/json"
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/nsqio/go-nsq"
	"time"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	packetCh := make(chan []byte, 1)
	packet, _ := utils.GetRandomPacket()
	data, _ := json.Marshal(packet)

	packetCh <- data

	for {
		select {
		case <-packetCh:
			for i := 0; i < 10000; i++ {
				go w.Publish("write_test", []byte("test"))
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
