package repository

import (
	"context"
)

type QueueRepository interface {
	Publish(ctx context.Context, topic string, message interface{}) error
	Subscribe(ctx context.Context, topic string, handler func(message []byte) error) error
	Close() error
}
