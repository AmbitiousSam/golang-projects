# Blog API
A RESTful API for managing blog posts and comments, built with Go, GORM, and SQLite.

---
## Table of Contents
- [Features](#features)
- [Technologies Used](#technologies-used)
- [API Endpoints](#api-endpoints)
  - [Post Routes](#post-routes)
  - [Comment Routes](#comment-routes)
- [Installation and Setup](#installation-and-setup)
- [Database](#database)
- [Testing the API](#testing-the-api)
  - [Posts](#posts)
  - [Comments](#comments)
- [Future Enhancements](#future-enhancements)

---
## Features
- Blog Posts:
  - Create, Read, Update, Delete (CRUD) operations.
  - Persistent storage using SQLite.
- Comments:
  - Add comments to blog posts.
  - Retrieve comments for a specific blog post.
- Relational Database:
  - Post and Comment models with a one-to-many relationship.
- Scalable Design:
  - Organized project structure for easy scalability and maintainability.
- Dynamic Routing:
  - Use of dynamic URL parameters to handle resources effectively.
---
## Technologies Used
- Programming Language: Go (Golang)
- Database: SQLite
- ORM: GORM
- Router: Gorilla Mux

---
## API Endpoints
### Post Routes
| Method | Endpoint | Description |
| --- | --- | --- |
| POST | /posts | Create a new blog post. |
| GET | /posts | Retrieve all blog posts. |
| PUT | /posts/{id} | Update a blog post by ID. |
| DELETE | /posts/{id} | Delete a blog post by ID. |
### Comment Routes
| Method | Endpoint | Description |
| --- | --- | --- |
| POST | /posts/{post_id}/comments | Add a comment to a specific blog post. |
| GET | /posts/{post_id}/comments | Retrieve comments for a specific post. |

---
## Installation and Setup
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd blog-api
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run main.go
   ```
   The server will start on http://localhost:8080.
## Database
- SQLite is used for persistent storage.
- GORM is used to manage database models and relationships.
- Tables:
  - Post:
    - ID (Primary Key)
    - Title (string)
    - Content (string)
  - Comment:
    - ID (Primary Key)
    - PostID (Foreign Key)
    - Content (string)

--- 
## Testing the API
### Posts
- Create a Post:
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"title": "My First Post", "content": "This is the content."}' http://localhost:8080/posts
  ```
- Get All Posts:
```bash
curl -X GET http://localhost:8080/posts
```
Update a Post:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Title", "content": "Updated content."}' http://localhost:8080/posts/1
```
Delete a Post:

```bash
curl -X DELETE http://localhost:8080/posts/1
```

### Comments
Add a Comment:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"post_id": 1, "content": "Great post!"}' http://localhost:8080/posts/1/comments
```

Get Comments for a Post:

```bash
curl -X GET http://localhost:8080/posts/1/comments
```

---

## Future Enhancements

**1. Authentication:**
Add JWT-based authentication for protected routes.

**2. Pagination:**
Implement pagination for posts and comments.

**3. Input Validation:**
Validate request payloads to prevent invalid data.

**4. Soft Deletes:**
Support soft deletion for posts and comments.

**5. Search and Filtering:**
Enable searching and filtering of posts by title or category.
