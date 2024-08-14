# Task Management API Documentation
## Postman doc Link: 
```sh
https://documenter.getpostman.com/view/33094138/2sA3s4kVvY
```
## Overview
The Task Management API provides endpoints for managing tasks with built-in authentication and authorization using JSON Web Tokens (JWT). The API supports user registration, login, and role-based access control, allowing only authenticated users to access certain endpoints. Admin users have additional privileges, such as the ability to create, update, and delete tasks, and promote other users to the admin role.

### Base URL
```
http://localhost:8080
```

## Prerequisites

### 1. Go
Ensure that you have [Go](https://golang.org/dl/) installed on your machine.

### 2. MongoDB
Make sure you have a MongoDB instance running. You can use a local MongoDB server or a cloud service like [MongoDB Atlas](https://www.mongodb.com/cloud/atlas).

### 3. Environment Variables
Before running the API, you need to set up environment variables for MongoDB and the JWT secret key.

### Environment Variables Setup

1. **Create a `.env` file** in the root of your project directory.

2. **Add the following variables** to the `.env` file:

    ```bash
    MONGODB_URI=mongodb://localhost:27017/taskmanager
    SECRET_KEY=your_secret_key_here
    ```

    - **MONGODB_URI:** This is the connection string to your MongoDB instance. Replace `mongodb://localhost:27017/taskmanager` with the appropriate URI if you're using MongoDB Atlas or a different setup.
    - **SECRET_KEY:** This is the secret key used to sign your JWT tokens. Replace `your_secret_key_here` with a strong, random string. Keep this key secure and do not share it publicly.

### Running the Code

1. **Install Dependencies**
   
   Ensure you have the necessary Go packages installed. You can use the `go mod tidy` command to install any missing dependencies:

   ```bash
   go mod tidy
   ```

2. **Run the API**

   To start the API server, use the following command:

   ```bash
   go run main.go
   ```

   This will start the API server on the default port (usually `localhost:8080`).

3. **Test the API**

   Once the server is running, you can test the API using Postman or any other API client. Make sure to include the JWT token in the `Authorization` header when accessing protected routes.

### Example .env File

Here's an example of what your `.env` file might look like:

```plaintext
MONGODB_URI=mongodb://localhost:27017/taskmanager
SECRET_KEY=supersecretkey123456
```

## Authentication and Authorization

### JSON Web Tokens (JWT)
This API uses JWT for authentication. After a successful login, a JWT token is generated, which must be included in the `Authorization` header of subsequent requests to access protected routes.

- **Token Type:** Bearer
- **Header Format:** `Authorization: Bearer <JWT_TOKEN>`

### User Roles
- **Admin:** Can create, update, delete tasks, and promote other users to admin.
- **User:** Can view tasks and access their details.

## Endpoints

### 1. User Registration

**Endpoint:** `POST /register`

**Description:** Create a new user account.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "StrongPassword123",
  "email": "john.doe@example.com",
  "phone": "1234567890"
}
```

**Response:**
- **201 Created**: User successfully registered.
- **400 Bad Request**: Validation errors or username already exists.

**Example:**
```json
{
  "message": "User registered successfully"
}
```

### 2. User Login

**Endpoint:** `POST /login`

**Description:** Authenticate user and generate JWT token.

**Request Body:**
```json
{
  "username": "john_doe",
  "password": "StrongPassword123"
}
```

**Response:**
- **200 OK**: Successfully authenticated. Returns JWT token.
- **401 Unauthorized**: Invalid username or password.

**Example:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Get All Tasks

**Endpoint:** `GET /tasks`

**Description:** Retrieve a list of all tasks. Accessible by both admin and regular users.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`

**Response:**
- **200 OK**: Returns the list of tasks.
- **401 Unauthorized**: Missing or invalid JWT token.

**Example:**
```json
[
  {
    "id": "60df5f4eaf4f8a001d706ccd",
    "title": "Complete API documentation",
    "description": "Write detailed API documentation for the task management system.",
    "status": "Pending",
    "created_at": "2024-08-01T12:00:00Z",
    "updated_at": "2024-08-01T12:00:00Z"
  },
  {
    "id": "60df5f4eaf4f8a001d706cce",
    "title": "Implement JWT authentication",
    "description": "Integrate JWT authentication into the task management API.",
    "status": "Completed",
    "created_at": "2024-08-02T12:00:00Z",
    "updated_at": "2024-08-02T12:00:00Z"
  }
]
```

### 4. Get Task by ID

**Endpoint:** `GET /tasks/:id`

**Description:** Retrieve a specific task by its ID. Accessible by both admin and regular users.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`

**Response:**
- **200 OK**: Returns the task details.
- **401 Unauthorized**: Missing or invalid JWT token.
- **404 Not Found**: Task not found.

**Example:**
```json
{
  "id": "60df5f4eaf4f8a001d706ccd",
  "title": "Complete API documentation",
  "description": "Write detailed API documentation for the task management system.",
  "status": "Pending",
  "created_at": "2024-08-01T12:00:00Z",
  "updated_at": "2024-08-01T12:00:00Z"
}
```

### 5. Create a Task (Admin Only)

**Endpoint:** `POST /tasks`

**Description:** Create a new task. Only accessible by admins.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`

**Request Body:**
```json
{
  "title": "Write test cases",
  "description": "Create and document test cases for the task management API.",
  "status": "Pending"
}
```

**Response:**
- **201 Created**: Task successfully created.
- **401 Unauthorized**: Missing or invalid JWT token.

**Example:**
```json
{
  "message": "Task created successfully",
  "task_id": "60df5f4eaf4f8a001d706ccf"
}
```

### 6. Update a Task (Admin Only)

**Endpoint:** `PUT /tasks/:id`

**Description:** Update an existing task. Only accessible by admins.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`

**Request Body:**
```json
{
  "title": "Update API documentation",
  "description": "Revise and update the API documentation to include recent changes.",
  "status": "In Progress"
}
```

**Response:**
- **200 OK**: Task successfully updated.
- **401 Unauthorized**: Missing or invalid JWT token.
- **404 Not Found**: Task not found.

**Example:**
```json
{
  "message": "Task updated successfully"
}
```

### 7. Delete a Task (Admin Only)

**Endpoint:** `DELETE /tasks/:id`

**Description:** Delete an existing task. Only accessible by admins.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`

**Response:**
- **200 OK**: Task successfully deleted.
- **401 Unauthorized**: Missing or invalid JWT token.
- **404 Not Found**: Task not found.

**Example:**
```json
{
  "message": "Task deleted successfully"
}
```

### 8. Promote a User to Admin (Admin Only)

**Endpoint:** `PUT /promote/:id`

**Description:** Promote a regular user to an admin. Only accessible by admins.

**Request Header:**
- **Authorization:** `Bearer <JWT_TOKEN>`


**Response:**
- **200 OK**: User successfully promoted to admin.
- **401 Unauthorized**: Missing or invalid JWT token.
- **404 Not Found**: User not found.

**Example:**
```json
{
  "message": "User promoted to admin successfully"
}
```
