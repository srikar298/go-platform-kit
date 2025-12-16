# cmd/

This directory contains the main entry points for different services or applications within the `go-platform-kit` project. Each subdirectory under `cmd/` represents a distinct executable application.

The goal is to keep the `main` function (and its immediate surrounding code) as lean as possible, primarily focusing on:
*   Parsing command-line arguments and configuration.
*   Initializing external dependencies (databases, caches, message brokers).
*   Wiring up the application layers (adapters, application services).
*   Starting the server or executing the command-line task.

This separation ensures that the core business logic (in `internal/`) remains reusable and testable, independent of the specific way the application is run.

Examples:
*   `cmd/server`: The entry point for the main HTTP API server.
*   `cmd/worker`: The entry point for a background worker service.
*   `cmd/cli`: The entry point for a command-line interface tool.
