package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type DumbKafkaConsumer struct {
	config *kafka.ConfigMap
	topic  string
}

func (dkc DumbKafkaConsumer) consume() {
	c, err := kafka.NewConsumer(dkc.config)

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{dkc.topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
