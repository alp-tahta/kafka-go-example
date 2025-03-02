package messaging

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaClient struct {
	producer sarama.AsyncProducer
	consumer sarama.Consumer
	brokers  []string
}

func NewKafkaClient(brokers []string, config *sarama.Config) (*KafkaClient, error) {
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		producer.Close()
		return nil, err
	}

	return &KafkaClient{
		producer: producer,
		consumer: consumer,
		brokers:  brokers,
	}, nil
}

func (k *KafkaClient) Publish(topic string, msg Message) error {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(msg.Key),
		Value: sarama.ByteEncoder(msg.Value),
	}

	k.producer.Input() <- kafkaMsg

	select {
	case success := <-k.producer.Successes():
		fmt.Printf("Message sent to partition %d at offset %d\n", success.Partition, success.Offset)
		return nil
	case err := <-k.producer.Errors():
		log.Printf("Failed to send message: %v", err)
		return err
	}

}

func (k *KafkaClient) Subscribe(topic string, handler func(msg Message)) error {
	partitionConsumer, err := k.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	go func() {
		for message := range partitionConsumer.Messages() {
			handler(Message{
				Key:   string(message.Key),
				Value: message.Value,
			})
		}
	}()

	return nil
}

func (k *KafkaClient) Close() error {
	if err := k.producer.Close(); err != nil {
		return err
	}
	return k.consumer.Close()
}
