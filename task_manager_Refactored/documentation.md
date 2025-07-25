# Task_Manager_Refactored

A simple RESTful Task Manager API built with Go, MongoDB, and Gin. It supports user authentication, task management (CRUD), and admin role promotion using Clean Architecture principles.

## Features

- User Sign Up & Login (JWT-based)
- Role-based access control (Admin/User)
- Create, Read, Update, Delete tasks
- Promote users to admin
- Clean project structure using Domain-Driven Design (DDD)

---

## Tech Stack

- **Go (Golang)**
- **Gin** – HTTP router
- **MongoDB** – Document store
- **JWT** – Token-based authentication
- **Clean Architecture** – Layered separation of concerns

---

## Project Layers

- `Delivery/` – HTTP handlers and route setup
- `Usecases/` – Business logic
- `Domain/` – Entities, interfaces, request/response 
- `Repositories/` – MongoDB logic
- `Infrastructure/` – JWT, auth middleware, password hashing

---

## Clean Architecture & Dependency Inversion

This project follows Clean Architecture principles, where:

- **High-level business rules** (Usecases) **do not depend on low-level details** (Database, HTTP).
- All **dependencies point inward**, toward the core logic.
- Interfaces in the `Domain` layer define contracts.
- Concrete implementations (e.g., MongoDB) live in `Repositories` and are injected into use cases.

This approach ensures the codebase is:

- **Modular** – Easy to maintain and extend
- **Testable** – Core logic has no framework dependency
- **Decoupled** – Logic and data access are isolated

---
