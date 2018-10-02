package main

import (
	"flag"
	"github.com/nats-io/go-nats"
	"github.com/sirupsen/logrus"
	"runtime"
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

	i := 0

	nc.Subscribe("benchmark", func(msg *nats.Msg) {
		i++
		nc.Publish(msg.Reply, []byte("reply"))
		if i%50000 == 0 {
			log.Info("TEST: ", i)
		}
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", "benchmark")

	runtime.Goexit()
}
