package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	db *redis.Client
}

func New(db *redis.Client) *RedisRepo {
	return &RedisRepo{db: db}
}

func (r *RedisRepo) AddLink(shortUrl, fullUrl string, c *context.Context) (string, error) {
	if !strings.HasPrefix(fullUrl, "http://") && !strings.HasPrefix(fullUrl, "https://") {
		fullUrl = "https://" + fullUrl
	}
	set, err := r.db.SetNX(*c, shortUrl, fullUrl, 0).Result()
	if err != nil {
		return "", err
	}
	if !set {
		return "", fmt.Errorf("short url already exists")
	}
	return shortUrl, nil
}

func (r *RedisRepo) GetById(id string, c *context.Context) (string, error) {
	return r.db.Get(*c, id).Result()
}
