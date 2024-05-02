package cmd

import (
	"os"

	"github.com/murilo-toddy/pixel/infrastructure/db"
	"github.com/murilo-toddy/pixel/app/kafka"
	"github.com/spf13/cobra"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start transactional consumer using Apache Kafka",
	Run: func(cmd *cobra.Command, args []string) {
        producer, err := kafka.NewKafkaProducer()
        database := db.Connect(os.Getenv("env"))
        if err != nil {
            panic(err)
        }
        deliveryChan := make(chan ckafka.Event)
        kafka.Publish("Hello", "test", producer, deliveryChan)
        go kafka.DeliveryReport(deliveryChan)

        kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChan)
        kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)
}
