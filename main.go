package main

import (
	"flag"
	"log"
	"time"

	"github.com/mauricioabreu/smart-monkey/message"
	"github.com/mauricioabreu/smart-monkey/service"
	"github.com/mauricioabreu/smart-monkey/store"
	"github.com/streadway/amqp"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "test-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "test-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "test-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "simple-consumer", "AMQP consumer tag (should not be blank)")
	lifetime     = flag.Duration("lifetime", 5*time.Second, "lifetime of process before shutdown (0s=infinite)")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("smart monkeys is starting...")

	c, err := message.NewConsumer(*uri, *exchange, *exchangeType, *queue, *bindingKey, *consumerTag)
	c.Consume(*queue, installHandler)
	if err != nil {
		log.Fatal(err)
	}

	if *lifetime > 0 {
		log.Printf("Running for %s", *lifetime)
		time.Sleep(*lifetime)
	} else {
		log.Printf("Running forever")
		select {}
	}

	log.Printf("Shutting down...")

	if err := c.Shutdown(); err != nil {
		log.Fatalf("Error during shutdown: %s", err)
	}
}

func installHandler(deliveries <-chan amqp.Delivery, done chan error) {
	repository := store.InMemoryStore()
	// Insert a new configuration in the storage
	key := "1"
	template := "foo"
	repository.StoreConfiguration(&store.Configuration{Key: key, Template: template})

	configurationService := service.InitService(repository)

	for d := range deliveries {
		log.Printf(
			"Got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		key := d.Body
		configurationService.Install(string(key))
		d.Ack(false)
	}
	log.Printf("Handle: deliveries channel closed")
	done <- nil
}
