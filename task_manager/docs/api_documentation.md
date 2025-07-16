<!-- Task Management API Documentation -->
    This is a simple Task Management REST API built using Go and the Gin Framework.
    It allows users to perform CRUD operations (Create, Read, Update, Delete) on tasks stored in an in-memory database.
    The API is designed for learning backend development with Go and does not persist data across restarts.
    All data is reset on server restart (in-memory storage only).
    All request/response payloads are in JSON format.

<!-- Base URL -->
    http://localhost:9090

<!-- Endpoints Summary -->
Method	    Endpoint	      Description
GET	        /tasks	        Get all tasks
GET	        /tasks/:id	    Get a task by ID
POST	    /tasks	        Create a new task
PUT	        /tasks/:id	    Update a task by ID
DELETE	    /tasks/:id	    Delete a task by ID

