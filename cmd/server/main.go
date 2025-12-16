package main

import (
	"context"
	"fmt"
	"go-platform-kit/configs"
	"go-platform-kit/internal/adapters"
	"go-platform-kit/internal/application"
	"log"
)

func main() {
	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("App Environment: %s\n", cfg.AppEnv)
	fmt.Printf("App Port: %d\n", cfg.AppPort)
	fmt.Printf("Log Level: %s\n", cfg.LogLevel)
	fmt.Printf("Timeout: %s\n", cfg.Timeout)

	// Initialize in-memory user repository (adapter)
	userRepo := adapters.NewInMemoryUserRepository()

	// Initialize user service (application layer)
	userService := application.NewUserService(userRepo)

	ctx := context.Background()

	// --- Example Usage ---

	// Create a new user
	fmt.Println("\nCreating a new user...")
	newUser, err := userService.CreateUser(ctx, "john_doe", "john.doe@example.com", "password123")
	if err != nil {
		log.Printf("Error creating user: %v", err)
	} else {
		fmt.Printf("User created: ID=%s, Username=%s, Email=%s\n", newUser.ID, newUser.Username, newUser.Email)
	}

	// Try to create the same user (should fail with "user with this email already exists")
	fmt.Println("\nAttempting to create user with existing email...")
	_, err = userService.CreateUser(ctx, "john_doe_2", "john.doe@example.com", "password123")
	if err != nil {
		log.Printf("Error creating user (expected): %v\n", err)
	} else {
		fmt.Println("Unexpected: User with existing email created successfully.")
	}

	// Get user by ID
	if newUser != nil {
		fmt.Printf("\nRetrieving user with ID: %s...\n", newUser.ID)
		foundUser, err := userService.GetUserByID(ctx, newUser.ID)
		if err != nil {
			log.Printf("Error getting user by ID: %v", err)
		} else {
			fmt.Printf("Found user: ID=%s, Username=%s, Email=%s\n", foundUser.ID, foundUser.Username, foundUser.Email)
		}
	}

	// Get non-existent user by ID
	fmt.Println("\nRetrieving non-existent user with ID: non-existent-id...")
	_, err = userService.GetUserByID(ctx, "non-existent-id")
	if err != nil {
		log.Printf("Error getting user (expected): %v\n", err)
	} else {
		fmt.Println("Unexpected: Non-existent user found.")
	}

	fmt.Println("\nApplication initialized and example operations performed.")
	fmt.Println("In a full application, a web server or other entry point would be started here.")
}
