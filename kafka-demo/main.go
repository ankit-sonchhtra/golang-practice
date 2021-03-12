package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic  = "kafka-test"
	broker = "localhost:9092"
)

func main() {
	ctx := context.Background()
	producer(ctx)
	consumer(ctx)
}

func producer(ctx context.Context) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	for i := 0; i < 10; i++ {
		err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("This is message: " + strconv.Itoa(i)),
		})

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Produced message with key: ", i)
		}
	}
}

func consumer(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})
	for {
		msg, err := reader.ReadMessage(ctx)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Consumed : ", string(msg.Value))
		}
	}
}
