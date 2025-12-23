# Task Manager

This project is a simple demonstration of a task manager with a Go backend and a basic HTML frontend.

The main goal of the project is to showcase the use of CI/CD with unit tests using GitHub Actions. It is intentionally kept simple and is not meant to be a full-featured production application.

## Purpose

- Demonstrate CI/CD pipelines with automated unit tests.
- Show continuous integration using GitHub Actions.
- Provide a minimal backend API and a simple frontend interface.
- Illustrate a clean and easy-to-test project setup.

## How to Run

### Backend (Go)
1. Navigate to the `backend/` directory.
2. Run in bash: `go run main.go`

example:
```bash
cd backend
go run main.go
```

### Frontend (HTML)
The frontend is a simple HTML and JavaScript interface.
It is served directly from the Go backend.

Open your browser at:
http://localhost:8080

### Tests
- To run unit tests: Go to `backend/` and run `go test`.
explample:

```bash     
cd backend
go test ./...   
```

## CI/CD
- A GitHub Actions workflow is defined in `.github/workflows/ci.yml`.
- The workflow runs automatically on every push and pull request to the main branch.
- It sets up Go and executes all backend unit tests to ensure code quality.