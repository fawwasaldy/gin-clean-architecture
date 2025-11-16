<div align="center">
  <h1>Gin Clean Architecture</h1>
  <p>
    A robust and scalable Go application built with the Gin framework, showcasing a practical implementation of Clean Architecture principles.
  </p>
  <p>
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go" alt="Go"></a>
    <a href="https://gin-gonic.com/"><img src="https://img.shields.io/badge/Gin-v1.11.0-007ACC?style=for-the-badge&logo=gin" alt="Gin"></a>
    <a href="https://gorm.io/"><img src="https://img.shields.io/badge/GORM-v1.31-9B59B6?style=for-the-badge&logo=gorm" alt="GORM"></a>
    <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL-18-336791?style=for-the-badge&logo=postgresql" alt="PostgreSQL"></a>
  </p>
</div>

---

## 🚀 About This Project

This repository serves as a blueprint for building maintainable, scalable, and testable web applications in Go. By strictly separating concerns, it demonstrates how **Clean Architecture** can be applied to create a system where business logic is independent of frameworks, databases, and UI, making it easier to evolve and test.

The core of the project is a user management system with JWT-based authentication, providing a solid foundation for more complex applications.

---

## 🏛️ Architecture: The Clean Approach

This project is structured around the principles of **Clean Architecture**, which organizes the code into independent layers with a strict dependency rule: outer layers can depend on inner layers, but not the other way around.

<div align="center">
  <img src="https://miro.medium.com/v2/resize:fit:640/format:webp/1*JWzL8VcHl13x0J5rDUZWzA.png" alt="Clean Architecture Diagram. Source: https://levelup.gitconnected.com" width="400"/>
</div>

-   **`Domain`**: The heart of the application. It contains the core business logic, entities (e.g., `User`, `Role`), and repository interfaces. It has zero dependencies on any other layer.
-   **`Application`**: Orchestrates the data flow. It contains the application-specific use cases (services) that execute the business logic. It depends only on the `Domain` layer.
-   **`Infrastructure`**: The outermost layer providing implementations for the interfaces defined in the inner layers. This includes database repositories (GORM), file storage adapters, and connections to other external services.
-   **`Presentation`**: The entry point for users. It handles HTTP requests, routing, and data presentation. In this project, it's implemented using the **Gin** framework and is responsible for translating HTTP requests into calls to the `Application` layer.

---

## ✨ Key Features

-   👤 **User Management**: Full CRUD operations for users.
-   🔐 **Secure Authentication**: JWT-based authentication with Access and Refresh Tokens.
-   🛡️ **Authorization**: Middleware for protecting routes and managing access control.
-   🗄️ **Database Integration**: Seamless data persistence with GORM and PostgreSQL.
-   📦 **Transactional Integrity**: Ensures data consistency for critical operations.
-   📄 **Query Logging**: A beautiful web interface to monitor and review database queries, organized by month.
-   🖼️ **File Uploads**: Handles user profile image uploads to local storage.

---

## 🛠️ Tech Stack

-   **Language**: Go
-   **Web Framework**: Gin
-   **Database**: PostgreSQL
-   **ORM**: GORM
-   **Authentication**: JWT (JSON Web Tokens)
-   **UUID**: Google UUID

---

## ⚙️ Dependency Injection with `samber/do`

This project uses the **samber/do** library for Dependency Injection (DI). DI helps us manage dependencies between components (like services, repositories, and controllers), making the code loosely coupled, more testable, and easier to maintain.

### When Should We Use DI?

You should use DI whenever one component needs a "service" provided by another component. Instead of the component creating its own dependency (e.g., `NewUserService` creating `NewUserRepository`), it should *receive* the dependency from a central "injector". This follows the **Inversion of Control (IoC)** principle.

### Where Are Dependencies Registered?

All dependency registrations are centralized in the `platform/provider/` directory.

-   **`platform/provider/provider.go`**: Registers global dependencies like the database (`*gorm.DB`), `JWTService`, and `TransactionRepository`. It also calls the modular providers.
-   **`platform/provider/adapter.go`**: Registers adapter implementations for domain ports (e.g., `port.FileStoragePort`).
-   **`platform/provider/user/provider.go`**: A feature-specific provider that registers all components related to the User feature (Controller, Service, Repository).

This entire registration process is initiated once in `main.go`.

### How to Use and Invoke Dependencies

#### 1. How to Register (Provide) a Dependency

To register a new service, add it to the appropriate provider file using `do.Provide`. The provider function should return the service interface and an error.

**Example (Registering `UserService`):**
```go
// in platform/provider/user/provider.go
func RegisterDependencies(injector do.Injector) {
    // ...
	do.Provide(injector, func(injector do.Injector) (service.UserService, error) {
		return service.NewUserService(injector), nil
	})
    // ...
}
```

#### 2. How to Use (Invoke) a Dependency

To get a dependency inside a component, add `do.Injector` as a parameter to its constructor (e.g., `NewUserService`). Then, use `do.MustInvoke[DependencyType](injector)` to retrieve the required service.

**Example (Inside NewUserService):**
```go
// in internal/application/service/user_service.go
func NewUserService(injector do.Injector) UserService {
    // Invoke dependencies from the injector 
	userRepository := do.MustInvoke[user.Repository](injector)
	refreshTokenRepository := do.MustInvoke[refresh_token.Repository](injector)
	userDomainService := do.MustInvoke[*user.Service](injector)
	jwtService := do.MustInvoke[JWTService](injector)

    // Return the struct populated with resolved dependencies 
	return &userService{
		userRepository:         userRepository, 
		refreshTokenRepository: refreshTokenRepository, 
		userDomainService:      userDomainService, 
		jwtService:             jwtService, 
		injector:               injector, // Store injector if needed for transactions 
	}
}
```

---

## 📂 Project Structure

The repository is organized to reflect the Clean Architecture layers and follows the standard Go project layout, making it easy to navigate and understand.

Core application code is placed within the `internal/` directory to enforce privacy and prevent external projects from importing it.

```
.
├── internal/
│   ├── application/
│   │   ├── request/
│   │   ├── response/
│   │   └── service/
│   ├── domain/
│   │   ├── identity/
│   │   ├── port/
│   │   ├── refresh_token/
│   │   ├── shared/
│   │   └── user/
│   ├── infrastructure/
│   │   ├── adapter/
│   │   └── database/
│   └── presentation/
│       ├── controller/
│       ├── message/
│       ├── middleware/
│       └── route/
├── command/
├── platform/
├── assets/
├── logs/
├── .env.example
├── go.mod
├── main.go
└── README.md
```

---

## 🚀 Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

-   [Go](https://golang.org/dl/) (version 1.25 or newer)
-   [PostgreSQL](https://www.postgresql.org/download/)
-   [Git](https://git-scm.com/downloads)

### Installation & Setup

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/fawwasaldy/gin-clean-architecture.git
    cd gin-clean-architecture
    ```

2.  **Configure Environment Variables:**
    Create a `.env` file in the root directory and populate it with your configuration. You can use the `.env.example` as a template.

    ```env
    APP_NAME=gin-clean-architecture
    IS_LOGGER=true
    
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=<your password>
    DB_NAME=<your database name>
    DB_PORT=5432
    
    NGINX_PORT=80
    GOLANG_PORT=8888
    APP_ENV=localhost
    
    JWT_SECRET=<your secret key>
    JWT_ISSUER=gin-clean-architecture
    JWT_ACCESS_EXPIRATION=15m
    JWT_REFRESH_EXPIRATION=7d

    AES_KEY=<your aes key>
    ```

3.  **Install Dependencies:**
    Go will automatically handle the installation of dependencies when you run the application.

4.  **Run Database Migrations:**
    Execute the following command to create the required tables in your database.
    ```bash
    go run main.go --migrate
    ```

5.  **Start the Server:**
    Use this command to launch the application.
    ```bash
    go run main.go
    ```
    The server will be live at `http://localhost:8888`.

---

## 🔌 API Endpoints

The following table lists the available API endpoints.

| Method   | Endpoint                  | Description                              | Authentication |
|:---------|:--------------------------|:-----------------------------------------|:--------------:|
| `POST`   | `/api/user/register`      | Register a new user                      |       No       |
| `POST`   | `/api/user/login`         | Log in to get an access token            |       No       |
| `POST`   | `/api/user/refresh-token` | Obtain a new access token                |       No       |
| `GET`    | `/api/user/me`            | Get the current user's profile           |      Yes       |
| `GET`    | `/api/user/`              | Get a paginated list of all users        |      Yes       |
| `PATCH`  | `/api/user/`              | Update the current user's profile        |      Yes       |
| `DELETE` | `/api/user/`              | Delete the current user's account        |      Yes       |
| `GET`    | `/logs`                   | View database query logs (current month) |       No       |
| `GET`    | `/logs/:month`            | View query logs for a specific month     |       No       |

---

## 🙏 Acknowledgements

This project is an enhancement and development of the foundation laid by **[go-gin-clean-starter](https://github.com/caknoooo/go-gin-clean-starter)**. A big thank you to **[Caknoooo](https://github.com/Caknoooo)** for creating an awesome starter template. We encourage you to visit and star the original repository as well!

---

## 🤝 Contributing

Contributions are welcome! If you have suggestions for improvements or find a bug, please feel free to fork the repository, make your changes, and submit a pull request. You can also open an issue with the "bug" or "enhancement" tag.

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
