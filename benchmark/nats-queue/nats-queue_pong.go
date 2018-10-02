package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"runtime"

	"github.com/nats-io/go-nats"
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

	nc.QueueSubscribe("benchmark", "benchmark-queue", func(msg *nats.Msg) {
		i++
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
