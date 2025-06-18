package rabbitmq

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const (
	_connString = "amqp://rabbit:rabbit@127.0.0.1:5672/"
	_queueName  = "test"
)

var (
	_messageData = uuid.MustParse("f86f517b-e509-42d4-a1f3-b8cc769cb938")

	_client   *Client
	_producer Producer
)

func TestOpenConn(t *testing.T) {
	t.Log("Open new RabbitMQ connection")

	var err error
	_client, err = NewClient(_connString)
	require.NoError(t, err, "open connection")

	t.Log("Connection was opened successfully!")
}

func TestCloseConn(t *testing.T) {
	t.Log("Close RabbitMQ connection")

	err := _client.Close()
	require.NoError(t, err, "close connection")
	t.Log("Connection was closed successfully!")

	t.Log("Open connection again (for other tests)")
	TestOpenConn(t)
}

func TestCreateProducer(t *testing.T) {
	t.Log("Create new RabbitMQ producer")

	var err error
	_producer, err = NewProducer(_client, _queueName)
	require.NoError(t, err, "init producer")
	t.Log("Producer was created successfully!")
}

func TestPublishText(t *testing.T) {
	t.Log("Publish text message")

	// use [:] to convert uuid to byte slice
	err := _producer.PublishText(_messageData[:])
	require.NoError(t, err, "publish message")
}
