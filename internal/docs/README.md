# internal/

This directory contains application-specific private code that should not be imported by other Go projects. It's a key component of structuring Go projects to enforce modularity and prevent accidental external dependencies.

Within `internal/`, the project adheres to the Hexagonal Architecture pattern, further segmenting responsibilities into:
*   `adapters/`: Implementations for interacting with external systems (databases, caches, third-party APIs).
*   `application/`: Business logic, use cases, and application services that orchestrate domain entities.
*   `domain/`: The core business entities, value objects, and rules that define the heart of the application.
*   `ports/`: Interfaces that define how the application (domain and application layers) interacts with external systems.

By keeping these core components within `internal/`, we ensure a clear separation of concerns and maintain a highly testable and interchangeable architecture.
