package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	Health() map[string]string
	SendMessage(channelName string, message any) error
	SubscribeToChannel(channelName string) <-chan *redis.Message
	GetPreviousMessagesFrom(channelName string, id int64) []map[string]any
}

type service struct {
	rdb *redis.Client
}

var (
	address  = os.Getenv("DB_ADDRESS")
	port     = os.Getenv("DB_PORT")
	password = os.Getenv("DB_PASSWORD")
	database = os.Getenv("DB_DATABASE")
)

func New() Service {
	num, err := strconv.Atoi(database)
	if err != nil {
		log.Fatalf(fmt.Sprintf("database incorrect %v", err))
	}

	fullAddress := fmt.Sprintf("%s:%s", address, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fullAddress,
		Password: password,
		DB:       num,
	})

	s := &service{rdb: rdb}

	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := s.rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	if result != "PONG" {
		log.Fatalf(fmt.Sprintf("Unexpected ping response: %s", result))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
