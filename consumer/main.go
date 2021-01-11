package main

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

const topic = "test-topic"

var consumerConfig = &kafka.ConfigMap{
	"bootstrap.servers": "localhost",
	"group.id":          "myGroup",
	"auto.offset.reset": "earliest",
}

func main() {
	dkc := DumbKafkaConsumer{
		config: consumerConfig,
		topic:  topic,
	}
	dkc.consume()
}
