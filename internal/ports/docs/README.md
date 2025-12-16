# internal/ports/

This directory defines the "ports" of the Hexagonal Architecture. Ports are interfaces that declare the contract for how the application's core logic (domain and application layers) interacts with external systems. They act as boundaries, ensuring the application remains decoupled from specific technology choices.

Responsibilities:
*   **Repository Interfaces:** Define methods for persisting and retrieving domain entities (e.g., `UserRepository`). These are typically driven ports (the application *needs* data).
*   **Service Interfaces:** Define external services the application needs to call (e.g., `EmailService`, `PaymentGateway`). These are also driven ports.
*   **Driven Ports:** Interfaces that the application *calls* to get something done (e.g., store data, send email).
*   **Driving Ports:** Interfaces that the application *exposes* for external actors to interact with (e.g., an `APIService` interface that a web adapter implements).

The `internal/ports/` directory ensures that the `internal/application/` and `internal/domain/` layers only depend on abstractions, allowing for flexible and interchangeable `internal/adapters/` implementations.
