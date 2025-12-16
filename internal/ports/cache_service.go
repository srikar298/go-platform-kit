package ports

import (
	"context"
	"time"
)

// CacheService defines the interface for a caching mechanism.
// It allows for storing and retrieving data from a cache.
type CacheService interface {
	// Get retrieves a value from the cache associated with the given key.
	// It returns the value and a boolean indicating if the key was found.
	Get(ctx context.Context, key string) (interface{}, bool)

	// Set stores a value in the cache with the given key and expiration duration.
	// It returns an error if the operation fails.
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error

	// Delete removes a key-value pair from the cache.
	Delete(ctx context.Context, key string) error
}
