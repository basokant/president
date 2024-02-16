package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (s *service) SendMessage(channelName string, message string) error {
	ctx := context.Background()

	return s.db.Publish(ctx, channelName, message).Err()
}

func (s *service) SubscribeToChannel(channelName string) <-chan *redis.Message {
	ctx := context.Background()

	pubsub := s.db.Subscribe(ctx, channelName)
	return pubsub.Channel()
}
