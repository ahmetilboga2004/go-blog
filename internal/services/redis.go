package services

import (
	"context"
	"time"

	"github.com/ahmetilboga2004/go-blog/internal/interfaces"
	"github.com/go-redis/redis/v8"
)

type redisService struct {
	client *redis.Client
}

func NewRedisService(addr, password string, db int) interfaces.RedisService {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisService{client: client}
}

func (s *redisService) BlacklistToken(token string, expiration time.Duration) error {
	ctx := context.Background()
	return s.client.Set(ctx, token, true, expiration).Err()
}

func (s *redisService) IsBlacklistedToken(token string) (bool, error) {
	ctx := context.Background()
	val, err := s.client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return val == "1", nil
}
