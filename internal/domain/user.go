package domain

import "time"

// User represents the core domain entity for a user.
type User struct {
	ID        string
	Username  string
	Email     string
	Password  string // Hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User instance.
func NewUser(id, username, email, hashedPassword string) (*User, error) {
	// In a real application, more robust validation would occur here.
	if id == "" || username == "" || email == "" || hashedPassword == "" {
		return nil, ErrInvalidUser
	}
	now := time.Now()
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// ErrInvalidUser is returned when a user cannot be created due to invalid data.
const ErrInvalidUser = domainError("invalid user data")

type domainError string

func (e domainError) Error() string {
	return string(e)
}
