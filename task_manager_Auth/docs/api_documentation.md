# Task Manager  Documentation

This project is a Task Management REST API built with Go, using the Gin framework for routing and middleware. MongoDB serves as the primary database to store user and task data. JWT (JSON Web Tokens) is used for secure, stateless authentication: users log in and receive a token that is included in subsequent requests. The token embeds the user's ID and role, which Gin middleware decodes to enforce role-based access control — allowing only Admins to create, update, delete tasks or promote users, while regular users can only view tasks.

## 1. API Security: JWT Authentication & Authorization

The API uses JSON Web Tokens (JWT) for secure authentication and role-based authorization.

### Authentication Flow

1.  **User Registration (`/register`):** New users create an account. The first registered user in an empty database is automatically assigned the `admin` role.
2.  **User Login (`/login`):** Authenticated users receive a JWT upon successful login.
3.  **Protected Routes:** Subsequent requests to protected endpoints must include this JWT in the `Authorization` header as a `Bearer Token`.
    * **Header Format:** `Authorization: Bearer <your_jwt_token>`

### Authorization (Role-Based Access)

* **Middleware:** JWT validation and role checking are handled by middleware.
* **User Roles:** `admin` and `regular user` roles are supported.
* **Access Control:**
    * **Admin-Only:** Routes for creating, updating, deleting tasks, promoting users (`/promote`), and viewing all users (`/admin/users`) are restricted to `admin` users only.
    * **Authenticated Users:** All users (admin and regular) can retrieve all tasks and retrieve tasks by ID.

### Secure Credentials

User passwords are securely stored using appropriate hashing techniques (e.g., bcrypt).
---

## 2. API Endpoints

**Base URL:** Typically `http://localhost:PORT` (where PORT is securely configured in the .env).

### Public Endpoints

These endpoints do not require authentication.

| Method | Endpoint        | Description                         |
| :----- | :-------------- | :---------------------------------- |
| `POST` | `/register`     | Creates a new user account.         |
| `POST` | `/login`        | Authenticates a user and returns a JWT. |

### Authenticated User Endpoints

These endpoints require a valid JWT in the `Authorization: Bearer` header.

| Method | Endpoint        | Description                               | Access Role |
| :----- | :-------------- | :---------------------------------------- | :---------- |
| `GET`  | `/tasks`        | Retrieves all tasks (for the authenticated user). | All Authenticated |
| `GET`  | `/tasks/:id`    | Retrieves a specific task by ID.          | All Authenticated |

### Admin-Only Endpoints

These endpoints require a valid JWT *and* the authenticated user must have the `admin` role.

| Method | Endpoint        | Description                               | Access Role |
| :----- | :-------------- | :---------------------------------------- | :---------- |
| `POST` | `/tasks`        | Creates a new task.                       | Admin       |
| `PUT`  | `/tasks/:id`    | Updates an existing task by ID.           | Admin       |
| `DELETE`| `/tasks/:id`   | Deletes a task by ID.                     | Admin       |
| `PUT`  | `/admin/promote`| Promotes a specified user to an administrator. | Admin       |
| `GET`  | `/admin/users`  | Retrieves a list of all users.            | Admin       |

---

## 3. Error Responses

The API returns standard HTTP status codes and a JSON body for errors.

| Status Code | Description                                   | Example JSON          |
| :---------- | :-------------------------------------------- | :-------------------- |
| `400 Bad Request` | Invalid input or missing data.                | `{"error": "Invalid payload."}` |
| `401 Unauthorized`| Missing or invalid JWT.                       | `{"error": "Unauthorized."}` |
| `403 Forbidden`   | Authenticated, but lacking required permissions. | `{"error": "Forbidden: Admin access required."}` |
| `404 Not Found`   | Resource not found.                           | `{"error": "Task not found."}` |
| `500 Internal Server Error` | Server-side issue.                        | `{"error": "An internal error occurred."}` |