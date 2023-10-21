package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func SetupProducerConnection(cfg ProducerConfig) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"client.id":         cfg.ClientID,
		"acks":              cfg.Ack,
	})
}

func Produce(prodConn *kafka.Producer, topic string, value []byte) error {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	// TODO check the feasibility of producing messages in goroutines
	err := prodConn.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)

	if err != nil {
		return errors.Wrap(err, "error while producing message")
	}

	chanOut := <-deliveryChan
	deliveryReport := chanOut.(*kafka.Message)
	return deliveryReport.TopicPartition.Error
}
