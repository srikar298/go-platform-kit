# internal/adapters/

This directory contains concrete implementations of the interfaces (ports) defined in `internal/ports/`. These implementations are the "adapters" in the Hexagonal Architecture, providing the necessary bridges between the application's core logic and external infrastructure concerns.

Responsibilities:
*   **Database Access:** Implementations for `UserRepository` (e.g., `PostgresUserRepository`, `InMemoryUserRepository`).
*   **Caching:** Implementations for `CacheService` (e.g., `RedisCacheService`, `InMemoryCacheService`).
*   **External API Clients:** Adapters for interacting with third-party APIs.
*   **Messaging:** Implementations for message brokers (e.g., Kafka, RabbitMQ).

Each adapter should only depend on the interfaces defined in `internal/ports/` and the specific external technology it integrates with. It should not directly depend on other internal application logic, ensuring that the core remains independent and testable.
