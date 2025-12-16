# internal/application/

This directory contains the application services and use cases that define the application's specific functionalities. These services orchestrate interactions between domain entities and various infrastructure components (via ports) to fulfill specific business requirements.

Responsibilities:
*   **Use Cases:** Implement specific actions or features of the application (e.g., `CreateUser`, `GetUserByID`, `ProcessOrder`).
*   **Orchestration:** Coordinate calls to domain entities and `ports` (like `UserRepository`, `CacheService`) to achieve a business goal.
*   **Business Logic:** Contains application-specific business rules that might not belong directly within a domain entity.
*   **Transaction Management:** Can manage transactions across multiple repository operations.
*   **Input/Output Handling:** Translating input from adapters (e.g., DTOs from HTTP handlers) into domain objects and translating domain objects back for output.

Application services should *not* contain infrastructure details. They depend on interfaces defined in `internal/ports/` and interact with `internal/domain/` entities.
