package main

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"strconv"
)

const topic = "test-topic"

var producerConfig = kafka.ConfigMap{
	"bootstrap.servers":  "localhost",
	"enable.idempotence": "true",
	"compression.type":   "snappy",
	"linger.ms":          "20",
}

func main() {
	dkp := DumbKafkaProducer{
		config: producerConfig,
		topic:  topic,
	}
	messages := generateMessages(1000000)

	dkp.Produce(messages)
}

func generateMessages(int) []string {
	var messages []string
	for i := 1; i < 1000000; i++ {
		messages = append(messages, strconv.Itoa(i))
	}

	return messages
}
