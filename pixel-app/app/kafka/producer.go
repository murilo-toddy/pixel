package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() (*ckafka.Producer, error) {
    configMap := &ckafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
    }
    producer, err := ckafka.NewProducer(configMap)
    if err != nil {
        return nil, err
    }
    return producer, nil
}

func Publish(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
    message := &ckafka.Message{
        TopicPartition: ckafka.TopicPartition{
            Topic: &topic,
            Partition: ckafka.PartitionAny,
        },
        Value: []byte(msg),
    }
    return producer.Produce(message, deliveryChan)
}

func DeliveryReport(deliveryChan chan ckafka.Event) {
    for event := range deliveryChan {
        switch ev := event.(type) {
        case *ckafka.Message:
            if ev.TopicPartition.Error != nil {
                fmt.Println("Could not deliver message:", ev.TopicPartition.Error)
            } else {
                fmt.Println("Message delivered to", ev.TopicPartition)
            }
        }
    }
}
