package database

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func (s *service) SendMessage(channelName string, message any) error {
	ctx := context.Background()

	err := s.rdb.Publish(ctx, channelName, message).Err()
	if err != nil {
		return err
	}

	return s.rdb.RPush(ctx, channelName, message).Err()
}

func (s *service) SubscribeToChannel(channelName string) <-chan *redis.Message {
	ctx := context.Background()
	return s.rdb.Subscribe(ctx, channelName).Channel()
}

func (s *service) GetPreviousMessagesFrom(channelName string, id int64) []map[string]any {
	ctx := context.Background()
	res := s.rdb.LRange(ctx, channelName, id, -1).Val()
	messages := make([]map[string]any, len(res))

	for i, m := range res {
		json.Unmarshal([]byte(m), &messages[i])
	}

	return messages
}
