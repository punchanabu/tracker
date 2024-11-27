package queue

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaQueue struct {
	writer  *kafka.Writer
	reader  map[string]*kafka.Reader
	brokers []string
}

func NewKafkaQueue(brokers []string) (*KafkaQueue, error) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		MaxAttempts:  3,
	}

	return &KafkaQueue{
		writer:  writer,
		reader:  make(map[string]*kafka.Reader),
		brokers: brokers,
	}, nil
}

func (k *KafkaQueue) Publish(ctx context.Context, topic string, message interface{}) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return k.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: value,
		Time:  time.Now(),
	})
}

func (k *KafkaQueue) Subscribe(ctx context.Context, topic string, handler func(message []byte) error) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:       k.brokers,
		Topic:         topic,
		GroupID:       "portfolio-tracker-group",
		StartOffset:   kafka.LastOffset,
		RetentionTime: 24 * time.Hour,
		MinBytes:      10e3, // 10KB
		MaxBytes:      10e6, // 10MB
	})

	k.reader[topic] = reader

	// Start process message with go routine
	go func() {
		defer reader.Close()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := reader.ReadMessage(ctx)
				/*
					Log Error but continue processing it to avoid panic
				*/
				if err != nil {
					continue
				}

				if err := handler(msg.Value); err != nil {
					continue
				}
			}
		}
	}()

	return nil
}

func (k *KafkaQueue) Close() error {
	if err := k.writer.Close(); err != nil {
		return err
	}

	for _, reader := range k.reader {
		if err := reader.Close(); err != nil {
			return err
		}
	}

	return nil
}