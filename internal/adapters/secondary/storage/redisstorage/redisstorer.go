package redisstorer

import (
	"context"
	"webpcdn/internal/ports"

	redis "github.com/go-redis/redis/v8"
)

func New() ports.Storer {
	store := &store{}
	store.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	store.ctx = context.Background()

	return store
}

type store struct {
	client *redis.Client
	ctx    context.Context
}

// isExists implements ports.Storer.
func (s *store) IsExists(filename string) bool {
	keyExists, err := s.client.Exists(s.ctx, filename).Result()
	if err != nil {
		return false
	}

	return keyExists == 1
}

// Read implements ports.Storer.
func (s *store) Read(filename string) ([]byte, error) {
	val, err := s.client.Get(s.ctx, filename).Result()
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

// Write implements ports.Storer.
func (s *store) Write(filename string, data []byte) error {
	err := s.client.Set(s.ctx, filename, data, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
