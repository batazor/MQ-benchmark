package main

//var (
//	log = logrus.New()
//)
//
//func main() {
//	decodeConfig := nsq.NewConfig()
//	c, err := nsq.NewConsumer("write_test", "test", decodeConfig)
//	if err != nil {
//		log.Panic("Could not create consumer")
//	}
//	//c.MaxInFlight defaults to 1
//
//	i := 0
//	c.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
//		if i%50000 == 0 {
//			log.Info("TEST: ", i)
//			i++
//		}
//
//		return nil
//	}), 10)
//
//	err = c.ConnectToNSQD("127.0.0.1:4150")
//	if err != nil {
//		log.Panic("Could not connect")
//	}
//
//	select {}
//}
