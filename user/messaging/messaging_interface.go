package messaging

type Message struct {
	Key     string
	Value   []byte
	Headers map[string]string
}

// MessagingClient defines an interface for message brokers like Kafka and RabbitMQ.
type MessagingClient interface {
	Publish(topic string, msg Message) error
	Subscribe(topic string, handler func(msg Message)) error
	Close() error
}
