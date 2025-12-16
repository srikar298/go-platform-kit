# cmd/server/

This directory contains the `main.go` file and other entry-point related code for the primary application server. It is responsible for:

*   **Application Startup:** Initializing the application.
*   **Dependency Injection (Wiring):** Assembling the various components (adapters, application services) and injecting their dependencies.
*   **Configuration Loading:** Loading application configuration from environment variables or files.
*   **Starting Infrastructure:** Starting HTTP servers, message consumers, etc.
*   **Graceful Shutdown:** Handling signals to ensure a clean shutdown of the application.

This `main` package should be kept minimal, focusing solely on the setup and orchestration, delegating business logic to the `internal/application` and `internal/domain` layers.
