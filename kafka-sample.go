package main

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func main() {

	bootstrapServers := "0.0.0.0:9092,0.0.0.0:19092"
	topic := "krc-adminclient-sampletopic"
	numParts := 1
	replicationFactor := 1

	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		return
	}
	defer a.Close()

	_, err = a.GetMetadata(&topic, false, 5000)
	if err != nil {
		fmt.Printf("Failed to connect to Kafka: %s\n", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	maxDur, err := time.ParseDuration("60s")
	if err != nil {
		fmt.Printf("Failed to parse duration: %s\n", err)
		return
	}
	topicConfig := kafka.TopicSpecification{
		Topic:             topic,
		NumPartitions:     numParts,
		ReplicationFactor: replicationFactor,
	}

	results, err := a.CreateTopics(ctx, []kafka.TopicSpecification{topicConfig}, kafka.SetAdminOperationTimeout(maxDur))
	if err != nil {
		fmt.Printf("Failed to create topic: %v\n", err)
		return
	}

	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}
