package adapters

import (
	"context"
	"go-platform-kit/internal/ports"
	"time"

	"github.com/patrickmn/go-cache"
)

var _ ports.CacheService = (*InMemoryCacheService)(nil)

// InMemoryCacheService implements ports.CacheService using github.com/patrickmn/go-cache.
type InMemoryCacheService struct {
	cache *cache.Cache
}

// NewInMemoryCacheService creates a new InMemoryCacheService with a default expiration and cleanup interval.
func NewInMemoryCacheService(defaultExpiration, cleanupInterval time.Duration) *InMemoryCacheService {
	c := cache.New(defaultExpiration, cleanupInterval)
	return &InMemoryCacheService{
		cache: c,
	}
}

// Get retrieves a value from the cache associated with the given key.
func (s *InMemoryCacheService) Get(ctx context.Context, key string) (interface{}, bool) {
	value, found := s.cache.Get(key)
	return value, found
}

// Set stores a value in the cache with the given key and expiration duration.
func (s *InMemoryCacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	s.cache.Set(key, value, expiration)
	return nil
}

// Delete removes a key-value pair from the cache.
func (s *InMemoryCacheService) Delete(ctx context.Context, key string) error {
	s.cache.Delete(key)
	return nil
}
