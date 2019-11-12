package main

import (
	"os"

	"github.com/cavargas40/kafka-go/consumer"
	"github.com/cavargas40/kafka-go/producer"
)

func main() {
	if os.Args[1] == "consumer" {
		consumer.StartConsumer()
	} else if os.Args[1] == "producer" {
		producer.StartProducer()
	}
}
