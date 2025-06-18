package rabbitmq

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	_connString  = "amqp://rabbit:rabbit@127.0.0.1:5672/"
	_queueName   = "test"
	_consumeTime = 20 * time.Second
)

var (
	_consumer Consumer
)

func TestMain(m *testing.M) {
	var err error
	// open connection
	client, err := NewClient(_connString)
	if err != nil {
		log.Fatalf("connect to RabbitMQ: %v", err)
	}
	// create consumer
	_consumer, err = NewConsumer(client, _queueName)
	if err != nil {
		log.Fatalf("init consumer: %v", err)
	}
	exitCode := m.Run()
	// close connection
	if err := client.Close(); err != nil {
		log.Fatalf("close connection with RabbitMQ: %v", err)
	}
	os.Exit(exitCode)
}

func TestConsume(t *testing.T) {
	t.Log("Consume messages")

	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	handler := func(msg amqp.Delivery) {
		t.Logf("Got message: %+v", msg)
	}

	// run consumer
	go func() {
		err := _consumer.Consume(ctx, handler)
		if err != nil {
			t.Error(err)
		}
		t.Log("Consumer was stopped successfully!")
		done <- struct{}{}
	}()

	// let consumer work for _consumeTime
	time.Sleep(_consumeTime)
	// send stop signal to consumer and wait for it
	cancel()
	<-done
}
