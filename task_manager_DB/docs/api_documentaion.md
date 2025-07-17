# 📝 Task Manager API Documentation

This project is a simple Task Manager backend built with **Go**, using the **Gin** web framework and **MongoDB** as the database. It exposes RESTful endpoints to manage tasks — users can create, read, update, and delete tasks.

The app is organized into a clean layered structure:
- `main.go` sets up the MongoDB connection and starts the HTTP server.
- `router/router.go` defines API routes and binds them to controller functions.
- `controllers/` handles incoming HTTP requests and delegates business logic.
- `service/` contains database logic and communicates with MongoDB.
- `models/` defines the Task schema using Go structs.

When a request is made (e.g., via Postman), the controller receives it, validates input, and calls the corresponding service function. The service layer runs MongoDB operations and returns results back to the controller, which sends an HTTP response.

This flow allows modular, testable, and maintainable code separation.
