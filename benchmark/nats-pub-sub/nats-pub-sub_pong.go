package main

//var (
//	log = logrus.New()
//)
//
//func main() {
//	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
//
//	nc, err := nats.Connect(*urls)
//	if err != nil {
//		log.Fatalf("Can't connect: %v\n", err)
//	}
//
//	i := 0
//	nc.Subscribe("benchmark", func(msg *nats.Msg) {
//		i += 1
//		if i%50000 == 0 {
//			log.Info("TEST: ", i)
//		}
//	})
//	nc.Flush()
//
//	if err := nc.LastError(); err != nil {
//		log.Fatal(err)
//	}
//
//	log.Printf("Listening on [%s]\n", "benchmark")
//
//	runtime.Goexit()
//}
