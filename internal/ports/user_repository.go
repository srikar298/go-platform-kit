package ports

import (
	"context"
	"go-platform-kit/internal/domain"
)

// UserRepository defines the methods for interacting with user data.
type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}
