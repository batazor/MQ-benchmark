package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/go-nats-streaming/pb"
)

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, m)
}

func main() {
	var showTime bool
	var startSeq uint64
	var startDelta string
	var deliverAll bool
	var deliverLast bool
	var durable string
	var qgroup string
	var unsubscribe bool

	//	defaultID := fmt.Sprintf("client.%s", nuid.Next())

	flag.BoolVar(&showTime, "t", false, "Display timestamps")
	// Subscription options
	flag.Uint64Var(&startSeq, "seq", 0, "Start at sequence no.")
	flag.BoolVar(&deliverAll, "all", false, "Deliver all")
	flag.BoolVar(&deliverLast, "last", false, "Start with last value")
	flag.StringVar(&startDelta, "since", "", "Deliver messages since specified time offset")
	flag.StringVar(&durable, "durable", "", "Durable subscriber name")
	flag.StringVar(&qgroup, "qgroup", "", "Queue group name")
	flag.BoolVar(&unsubscribe, "unsubscribe", false, "Unsubscribe the durable on exit")

	log.SetFlags(0)
	flag.Parse()

	sc, err := stan.Connect("benchmark", "benchmark", stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", stan.DefaultNatsURL, "benchmark", "benchmark")

	subj, i := "benchmark", 0

	mcb := func(msg *stan.Msg) {
		i++
		printMsg(msg, i)
	}

	startOpt := stan.StartAt(pb.StartPosition_NewOnly)

	if startSeq != 0 {
		startOpt = stan.StartAtSequence(startSeq)
	} else if deliverLast {
		startOpt = stan.StartWithLastReceived()
	} else if deliverAll {
		log.Print("subscribing with DeliverAllAvailable")
		startOpt = stan.DeliverAllAvailable()
	} else if startDelta != "" {
		ago, err := time.ParseDuration(startDelta)
		if err != nil {
			sc.Close()
			log.Fatal(err)
		}
		startOpt = stan.StartAtTimeDelta(ago)
	}

	sub, err := sc.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	if err != nil {
		sc.Close()
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], clientID=[%s], qgroup=[%s] durable=[%s]\n", subj, "benchmark", qgroup, durable)

	if showTime {
		log.SetFlags(log.LstdFlags)
	}

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			// Do not unsubscribe a durable on exit, except if asked to.
			if durable == "" || unsubscribe {
				sub.Unsubscribe()
			}
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
