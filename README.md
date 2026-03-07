# Todo API

A simple RESTful API for managing todos built with Go and the Gin framework. This project provides user authentication via JWT tokens and in-memory storage for users and tasks.

## Features

- **User Registration and Login**: Secure authentication with JWT tokens.
- **Profile Management**: View and update user profiles.
- **Todo Management**: Create, list, update, delete, and fulfill todos.
- **Protected Routes**: All API endpoints require authentication.
- **In-Memory Storage**: Lightweight store for development and testing.

## Pros

- **Simplicity**: Easy to set up and run without external dependencies.
- **Fast Development**: Quick prototyping with in-memory data.
- **Lightweight**: Minimal resource usage for small-scale applications.
- **Secure Auth**: JWT-based authentication ensures protected access.

## Cons

- **No Persistence**: Data is lost on restart (in-memory only).
- **Not Scalable**: In-memory store isn't suitable for production or multi-user loads.
- **Basic Logging**: Limited logging and monitoring features.
- **No Database**: Lacks advanced querying and data integrity.

## Tools Used

- **Go**: Programming language for backend logic.
- **Gin**: Web framework for HTTP routing and middleware.
- **JWT**: Library for token-based authentication (`github.com/golang-jwt/jwt/v5`).

## Getting Started

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Execute `go run main.go` to start the server on `localhost:8080`.
4. Use tools like Postman to test endpoints (e.g., register, login, manage todos).

For detailed API documentation, refer to the code comments in the handlers.