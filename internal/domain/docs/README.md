# internal/domain/

This directory is the heart of the application, representing the core business logic and entities. It defines what the system *is* and *does* at an abstract level, independent of any external concerns. This is where Domain-Driven Design (DDD) principles are most heavily applied.

Responsibilities:
*   **Entities:** Core objects with a distinct identity that have a lifecycle and encapsulate behavior (e.g., `User`, `Product`, `Order`).
*   **Value Objects:** Objects that represent descriptive aspects of the domain with no conceptual identity (e.g., `EmailAddress`, `Money`).
*   **Aggregates:** Clusters of domain objects that are treated as a single unit for data changes, ensuring consistency.
*   **Domain Services:** Operations that don't naturally fit into an entity or value object, often coordinating multiple entities.
*   **Domain Events:** Events that signify something important happened in the domain.
*   **Business Rules:** Encapsulation of the core business logic and invariants.

The `domain/` layer should have no dependencies on `application/`, `adapters/`, `ports/`, or `cmd/`. It is the most inner layer of the Hexagonal Architecture, purely focused on solving the business problem.
