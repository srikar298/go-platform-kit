package application

import (
	"context"
	"errors"
	"go-platform-kit/internal/domain"
	"go-platform-kit/internal/ports"
	"time"

	"github.com/google/uuid"
)

// UserService manages user-related business logic.
type UserService struct {
	userRepo    ports.UserRepository
	cacheService ports.CacheService // Add cache service field
}

// NewUserService creates a new UserService.
func NewUserService(repo ports.UserRepository, cacheService ports.CacheService) *UserService {
	return &UserService{
		userRepo:    repo,
		cacheService: cacheService, // Assign cache service
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, username, email, password string) (*domain.User, error) {
	// In a real application, password would be hashed here.
	// For simplicity, we'll store it as is for now (NOT PRODUCTION SAFE).
	hashedPassword := password // TODO: Implement proper password hashing

	id := uuid.New().String()
	user, err := domain.NewUser(id, username, email, hashedPassword)
	if err != nil {
		return nil, err
	}

	// Check if a user with this email already exists
	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}
	// Handle error if it's not a "not found" error
	if err != nil && err.Error() != "user not found" {
		return nil, err
	}

	if err := s.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}
	// Invalidate cache for the user if created
	s.cacheService.Delete(ctx, user.ID)
	return user, nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	// Try to get from cache first
	if cachedUser, found := s.cacheService.Get(ctx, id); found {
		if user, ok := cachedUser.(*domain.User); ok {
			return user, nil
		}
	}

	// If not in cache, get from repository
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Store in cache
	// Use a reasonable expiration time, e.g., 5 minutes.
	s.cacheService.Set(ctx, id, user, 5*time.Minute)
	return user, nil
}

// UpdateUser updates an existing user's information.
func (s *UserService) UpdateUser(ctx context.Context, id, username, email string) (*domain.User, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email
	user.UpdatedAt = time.Now()

	// In a real application, you might want to handle password changes separately
	// or ensure email uniqueness validation here if email is changed.

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}
	// Invalidate cache after update
	s.cacheService.Delete(ctx, id)
	return user, nil
}

// DeleteUser deletes a user by their ID.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	// Invalidate cache after deletion
	s.cacheService.Delete(ctx, id)
	return nil
}
