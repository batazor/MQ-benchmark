package amqp

import (
	"github.com/batazor/MQ-benchmark/utils"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()

	AMQP_API           = utils.Getenv("AMQP_API", "amqp://guest:guest@localhost:5672/")
	AMQP_NAME_QUEUE    = utils.Getenv("AMQP_NAME_QUEUE", "benchmark")
	AMQP_BINDING_KEY   = utils.Getenv("AMQP_BINDING_KEY", "")
	AMQP_CONSUMER_TAG  = utils.Getenv("AMQP_CONSUMER_TAG", "")
	AMQP_EXCHANGE_LIST = utils.Getenv("AMQP_EXCHANGE_LIST", "benchmark")
	AMQP_EXCHANGE_TYPE = utils.Getenv("AMQP_EXCHANGE_TYPE", "headers")

	CONSUMER = &Consumer{}
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Listen() Consumer {
	CONSUMER = NewConsumer(AMQP_API, AMQP_EXCHANGE_LIST, AMQP_EXCHANGE_TYPE, AMQP_NAME_QUEUE, AMQP_BINDING_KEY, AMQP_CONSUMER_TAG)

	err := CONSUMER.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Run AMQP")

	return *CONSUMER
}
