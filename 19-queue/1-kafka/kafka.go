package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// Client - клиент очереди Kafka.
type Client struct {
	Reader *kafka.Reader
	Writer *kafka.Writer
}

// New создает и инициализирует клиента Kafka.
func New(brokers []string, topic string, groupId string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" || groupId == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}

	c := Client{}

	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	c.Writer = &kafka.Writer{
		Addr:                   kafka.TCP(brokers[0]),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &c, nil
}

func main() {
	// Инициализация клиента Kafka.
	kfk, err := New(
		[]string{"localhost:29092"},
		"test-topic",
		"test-consumer-group",
	)
	if err != nil {
		log.Fatal(err)
	}

	go kfk.producer()
	kfk.consumer()
}

func (c *Client) producer() {
	for {
		msg := kafka.Message{
			Value: []byte(time.Now().String()),
		}
		err := c.Writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}

func (c *Client) consumer() {
	for {
		msg, err := c.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println(string(msg.Value))

		err = c.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}
