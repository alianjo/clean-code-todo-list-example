# Go To-Do Clean SOLID

A simple To-Do List API built in Go, showcasing **clean architecture**, **SOLID principles**, and a design focused on minimizing the **cost of change**. This project integrates with MongoDB to manage users and their tasks, providing a foundation for extensible and maintainable software.

## Features
- **User Creation**: Register users via a POST endpoint.
- **Task Management**: Create and retrieve tasks for users.
- **Clean Code**: Readable, modular, and well-documented code.
- **SOLID Principles**: Applied to ensure flexibility and scalability.
- **MongoDB Integration**: Persistent storage for users and tasks.

## Tech Stack
- **Golang**: Backend language.
- **MongoDB**: Database for persistence.
- **Gorilla Mux**: HTTP routing.
- **Docker**: Containerization for easy setup.

## Project Structure
The project follows a clean architecture pattern to separate concerns:
```
go-todo-clean-solid/
├── domain/         # Core entities and interfaces
├── usecase/        # Business logic
├── interfaces/     # HTTP handlers
├── infrastructure/ # External services (e.g., MongoDB)
├── Dockerfile      # Docker configuration for the app
├── docker-compose.yml  # Docker Compose setup
└── main.go         # Application entry point
```

## Getting Started

### Prerequisites
- **Docker** and **Docker Compose** installed (see [Docker Install](https://docs.docker.com/get-docker/)).
- Alternatively, for local setup without Docker:
  - Go 1.21+ installed.
  - MongoDB running locally (`mongodb://localhost:27017`) or a MongoDB Atlas URI.
  - Git installed.

### Setting Up with Docker
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/[your-username]/go-todo-clean-solid.git
   cd go-todo-clean-solid
   ```
2. **Start the Services**:
   - Run Docker Compose to build and start both the app and MongoDB:
     ```bash
     docker-compose up --build
     ```
   - This starts MongoDB on `mongodb://localhost:27017` and the API on `http://localhost:8080`.
3. **Stop the Services** (when done):
   - Press `Ctrl+C` in the terminal, then:
     ```bash
     docker-compose down
     ```
   - To remove volumes (data): `docker-compose down -v`.

### Setting Up Without Docker (Local MongoDB)
#### Install MongoDB Locally
1. **macOS**:
   ```bash
   brew install mongodb-community
   brew services start mongodb-community
   ```
2. **Ubuntu**:
   ```bash
   sudo apt-get update
   sudo apt-get install -y mongodb
   sudo systemctl start mongodb
   ```
3. **Windows**:
   - Download from [MongoDB](https://www.mongodb.com/try/download/community), install, and run `mongod`.

#### Install and Run the App
1. Clone the repository (as above).
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the app:
   ```bash
   go run main.go
   ```

### API Endpoints
- **POST /users**  
  Create a user.  
  Request: `{"name": "Alice", "email": "alice@example.com"}`  
  Response: `{"id": "user-123", "name": "Alice", "email": "alice@example.com"}`

- **POST /tasks**  
  Create a task.  
  Request: `{"user_id": "user-123", "description": "Finish project"}`  
  Response: `{"id": "task-456", "user_id": "user-123", "description": "Finish project", "done": false}`

- **GET /tasks/{userID}**  
  Get tasks for a user.  
  Response: `[{"id": "task-456", "user_id": "user-123", "description": "Finish project", "done": false}]`

### Example Usage
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com"}' http://localhost:8080/users
curl -X POST -H "Content-Type: application/json" -d '{"user_id":"user-123","description":"Finish project"}' http://localhost:8080/tasks
curl http://localhost:8080/tasks/user-123
```

## Design Philosophy
This project emphasizes:
- **Clean Code**: Meaningful names, small functions, and minimal comments where code is self-explanatory.
- **SOLID Principles**:
  - *Single Responsibility*: Each layer/module has one job.
  - *Open/Closed*: Extensible via new implementations (e.g., swap MongoDB for PostgreSQL).
  - *Liskov Substitution*: Interfaces ensure substitutability.
  - *Interface Segregation*: Small, focused interfaces.
  - *Dependency Inversion*: High-level modules depend on abstractions.
- **Cost of Change**: Modular design reduces the effort to modify or extend features (e.g., adding authentication or new endpoints).

## Why This Matters
As a software engineer, I prioritize maintainability and scalability. This project demonstrates how clean architecture and SOLID principles reduce technical debt and make future changes cost-effective.

## Future Enhancements
- Add authentication (JWT).
- Support task updates and deletions.
- Use UUIDs for IDs.
- Add unit tests.

## Contributing
Feel free to fork this repo, submit issues, or send pull requests! Contributions are welcome.

## License
MIT License - see [LICENSE](LICENSE) for details.