# Caching in Hexagonal Architecture

## What is Caching?

Caching is a technique used to store copies of frequently accessed data in a temporary storage area (the "cache") so that future requests for that data can be served faster. The primary goal of caching is to improve application performance by reducing the latency of data retrieval and decreasing the load on slower, more resource-intensive data sources (like databases or external APIs).

**Benefits of Caching:**
*   **Improved Performance:** Faster response times for users.
*   **Reduced Database/API Load:** Less strain on backend systems, leading to better scalability and stability.
*   **Lower Costs:** Can reduce infrastructure costs by minimizing database queries or external API calls.

## Types of Caching

Caching can be implemented at various levels of an application and infrastructure:

1.  **In-Memory Cache:** Data is stored directly in the application's RAM.
    *   **Pros:** Extremely fast access.
    *   **Cons:** Data is lost if the application restarts; limited by server memory; not shared across multiple application instances.
    *   **Use cases:** Small, frequently accessed datasets; application-specific configurations; session data in single-instance apps.

2.  **Distributed Cache:** Data is stored in a separate, dedicated caching server or cluster (e.g., Redis, Memcached).
    *   **Pros:** Shared across multiple application instances; data persists even if an app instance restarts; scalable independently.
    *   **Cons:** Network latency introduces some overhead; adds another component to manage.
    *   **Use cases:** Session management in clustered applications; shared product catalogs; rate limiting.

3.  **Client-Side Cache (Browser/CDN):** Data cached by the user's browser or a Content Delivery Network.
    *   **Pros:** Closest to the user, fastest response; offloads traffic from application servers.
    *   **Cons:** Complex invalidation strategies; data consistency challenges.
    *   **Use cases:** Static assets (images, CSS, JS); public API responses.

4.  **Database Cache:** Some databases offer built-in caching mechanisms.

## Caching in Go

Go provides excellent primitives for implementing in-memory caches (maps, sync.RWMutex). For more advanced features like eviction policies (LRU, LFU) or time-based expiration, external libraries like `ristretto` or `go-cache` are popular choices. For distributed caching, client libraries for Redis (`go-redis`) or Memcached are widely used.

## Caching in Hexagonal Architecture

In a Hexagonal Architecture (Ports and Adapters), caching is an **infrastructure concern**. This means the core domain and application logic should not directly know about or depend on the specific caching technology being used. Instead, caching should be treated as an **adapter** that implements a **port**.

Here's how it fits:

1.  **Port (Interface) Definition:**
    *   In the `internal/ports` package, we define a generic interface, say `CacheService`. This interface specifies *what* a cache can do (e.g., `Get(key string) (interface{}, bool)`, `Set(key string, value interface{}, duration time.Duration)`).
    *   The `application` layer (e.g., `UserService`) depends on this `CacheService` interface. It expresses its *need* for caching without knowing *how* it's implemented.

2.  **Adapter Implementation:**
    *   In the `internal/adapters` package, we create concrete implementations of the `CacheService` interface. Examples:
        *   `InMemoryCacheAdapter`: Uses a Go map for in-memory caching.
        *   `RedisCacheAdapter`: Uses the `go-redis` library to connect to a Redis instance.
    *   These adapters handle the technical details of interacting with the specific caching technology.

3.  **Application Layer (e.g., `UserService`) Usage:**
    *   The `UserService` receives an instance of `ports.CacheService` through dependency injection (e.g., in its constructor).
    *   When the `UserService` needs to retrieve data, it first checks the `CacheService`. If the data is found and valid, it returns the cached data.
    *   If not found in the cache, it fetches the data from the primary source (e.g., `UserRepository`), then stores it in the `CacheService` before returning it.
    *   The `UserService` remains decoupled from `InMemoryCacheAdapter` or `RedisCacheAdapter`; it only knows about `ports.CacheService`.

**Benefits of this approach:**
*   **Decoupling:** The core application logic is free from caching implementation details.
*   **Testability:** You can easily mock the `CacheService` interface for testing the `UserService` without needing a real cache.
*   **Flexibility:** You can swap out the caching technology (e.g., switch from in-memory to Redis) by simply providing a different adapter implementation at application startup, without changing the `UserService` code.
*   **Clarity:** The role of caching is clearly defined as an infrastructure concern.

## Example Scenario: Caching User Retrieval

Consider our `UserService` which fetches users via `FindByID` or `FindByEmail` from the `UserRepository`. This operation might be slow if the `UserRepository` interacts with a database.

With caching, the flow in `UserService.GetUserByID` would look like this:

1.  Receive `userID`.
2.  Check `CacheService` for `userID`.
3.  If `userID` is found in cache: return cached `User`.
4.  If `userID` is *not* found in cache:
    a.  Call `UserRepository.FindByID(userID)`.
    b.  If user is found: Store `User` in `CacheService` with an appropriate expiration.
    c.  Return `User`.
5.  If `UserRepository` returns an error (e.g., user not found), do not cache.

This ensures that frequently requested users are served quickly from the cache, reducing database load.
