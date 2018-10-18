package message

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

// NewConsumer : create a new consumer to start processing all incoming messages
func NewConsumer(amqpURI, exchange, exchangeType, queueName, key, ctag string) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),
	}
	var err error

	log.Printf("Dialing %q", amqpURI)
	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}

	go func() {
		fmt.Printf("Closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("Getting channel...")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("Declaring exchange (%q)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("Exchange declare: %s", err)
	}

	log.Printf("Declared exchange, declaring queue %q", queueName)
	queue, err := c.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Queue declare: %s", err)
	}

	log.Printf("Declared queue (%q %d messages, %d consumers), binding to exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, key)

	if err = c.channel.QueueBind(
		queue.Name,
		key,
		exchange,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("Queue bind: %s", err)
	}

	return c, nil
}

type callback func(deliveries <-chan amqp.Delivery, done chan error)

func (c *Consumer) Consume(queueName string, hn callback) error {
	log.Printf("Queue bound to exchange, consuming (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queueName,
		c.tag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Queue consume: %s", err)
	}

	go hn(deliveries, c.done)
	return nil
}

func (c *Consumer) Shutdown() error {
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("Consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer log.Printf("AQMP shutdown ok")

	return <-c.done
}
