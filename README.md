<div align="center">
  <h1>Gin Clean Architecture</h1>
  <p>
    A robust and scalable Go application built with the Gin framework, showcasing a practical implementation of Clean Architecture principles.
  </p>
  <p>
    <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go" alt="Go"></a>
    <a href="https://gin-gonic.com/"><img src="https://img.shields.io/badge/Gin-v1.10.1-007ACC?style=for-the-badge&logo=gin" alt="Gin"></a>
    <a href="https://gorm.io/"><img src="https://img.shields.io/badge/GORM-v1.30-9B59B6?style=for-the-badge&logo=gorm" alt="GORM"></a>
    <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL-16-336791?style=for-the-badge&logo=postgresql" alt="PostgreSQL"></a>
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
-   ⚙️ **Automatic Migrations**: Keep your database schema in sync with your models automatically.
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

## 📂 Project Structure

The repository is organized to reflect the Clean Architecture layers, making it easy to navigate and understand.

```
.
├── 📁 application/      # Application Layer: Use Cases/Services
│   ├── request/
│   └── response/
│   └── service/
├── 📁 command/           # CLI commands (e.g., migrations)
├── 📁 domain/            # Domain Layer: Entities & Core Logic
│   ├── identity/
│   ├── port/
│   ├── refresh_token/
│   ├── shared/
│   └── user/
├── 📁 infrastructure/    # Infrastructure Layer: DB, Adapters
│   ├── adapter/
│   └── database/
├── 📁 platform/          # Shared platform utilities (e.g., pagination)
├── 📁 presentation/      # Presentation Layer: Controllers, Routes
│   ├── controller/
│   ├── message/
│   ├── middleware/
│   └── route/
├── assets/              # Static assets (e.g., uploaded images)
├── logs/                # Log files
├── .env.example         # Example environment variables
├── go.mod               # Go module dependencies
├── main.go              # Application entry point
└── README.md
```

---

## 🚀 Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

-   [Go](https://golang.org/dl/) (version 1.24 or newer)
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
    # Application Settings
    APP_ENV=development
    GOLANG_PORT=8888
    IS_LOGGER=true

    # Database Connection
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASS=your_db_password
    DB_NAME=your_db_name

    # JWT Configuration
    JWT_SECRET=a-very-strong-and-secret-key
    JWT_ISSUER=my-app
    JWT_ACCESS_EXPIRATION=15m
    JWT_REFRESH_EXPIRATION=7d
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
    go run main.go --run
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

## 🤝 Contributing

Contributions are welcome! If you have suggestions for improvements or find a bug, please feel free to fork the repository, make your changes, and submit a pull request. You can also open an issue with the "bug" or "enhancement" tag.

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

