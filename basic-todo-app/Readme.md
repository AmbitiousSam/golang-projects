## Project: Todo List API (Beginner)
### What We Learned
#### Dynamic Data Management:

Introduced an in-memory data store (todos slice) to persist data between requests.
Managed state within the application.

1. **CRUD Operations**:
Implemented core backend functionality:
```
POST: Add a new task.
GET: Retrieve all tasks.
PUT: Update an existing task.
DELETE: Remove a specific task.
```

2. **Working with JSON**:
Decoding: Converted JSON request body to Go structs (json.NewDecoder).
Encoding: Serialized Go structs into JSON responses (json.NewEncoder).

3. **HTTP Methods**:
Used different HTTP methods (GET, POST, PUT, DELETE) to define route-specific actions.

4. **Error Handling**:
Validated input data and returned meaningful HTTP error codes (400, 404, etc.).
Handled scenarios where a task doesn’t exist or input is invalid.

5. **Routing with HTTP Verbs**:
Handled multiple HTTP methods on the same route (/todos) using r.Method.

6. **State Management**:
Managed a dynamic list of tasks in memory that persists for the server's runtime.

#### New Concepts Implemented

1. **In-Memory Data Store**:
Used a slice to simulate database storage for todos.
2. **Structured Responses**:
Responded with meaningful and structured JSON objects.
3. **Input Validation**:
Ensured request body integrity before processing.
4. **RESTful API Design**:
Followed REST principles for designing API endpoints.
