package adapters

import (
	"context"
	"errors"
	"go-platform-kit/internal/domain"
	"go-platform-kit/internal/ports"
	"sync"
)

var _ ports.UserRepository = (*InMemoryUserRepository)(nil)

// InMemoryUserRepository implements ports.UserRepository for in-memory storage.
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewInMemoryUserRepository creates a new InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

// Save saves a user to in-memory storage.
func (r *InMemoryUserRepository) Save(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		// In a real scenario, you might want to return an error like ErrUserAlreadyExists
		// or handle updates separately. For simplicity, we'll just overwrite.
	}
	r.users[user.ID] = user
	return nil
}

// FindByID retrieves a user by ID from in-memory storage.
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found") // Or a custom domain error
	}
	return user, nil
}

// FindByEmail retrieves a user by email from in-memory storage.
func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found") // Or a custom domain error
}

// Update updates an existing user in in-memory storage.
func (r *InMemoryUserRepository) Update(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found for update") // Or a custom domain error
	}
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID from in-memory storage.
func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found for deletion") // Or a custom domain error
	}
	delete(r.users, id)
	return nil
}
