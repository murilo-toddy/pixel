package kafka

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

type KafkaProcessor struct {
    Database *gorm.DB
    Producer *ckafka.Producer
    DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(
    database *gorm.DB, 
    producer *ckafka.Producer,
    deliveryChan chan ckafka.Event,
) *KafkaProcessor {
    return &KafkaProcessor{
        Database: database,
        Producer: producer,
        DeliveryChan: deliveryChan,
    }
}

func (k *KafkaProcessor) Consume() {
    configMap := &ckafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
        "group.id": "consumergroup",
        "auto.offset.reset": "earliest",
    }
    consumer, err := ckafka.NewConsumer(configMap)
    if err != nil {
        panic(err)
    }

    topics := []string{"test"}
    consumer.SubscribeTopics(topics, nil)
    fmt.Println("Kafka consumer setup")

    for {
        if msg, err := consumer.ReadMessage(-1); err == nil {
            k.processMessage(msg)
        }
    }
}

func (k *KafkaProcessor) processMessage(msg *ckafka.Message) {
    transactionTopic := "transactions"
    transactionConfirmationTopic := "transaction_confirmation"

    switch topic := *msg.TopicPartition.Topic; topic {
    case transactionTopic:
        k.processTransaction(msg)
    case transactionConfirmationTopic:
    default:
        fmt.Println("Got message from invalid topic:", topic, string(msg.Value))
    }
}

func (k *KafkaProcessor) processTransaction(msg *ckafka.Message) error {
    return nil
}

