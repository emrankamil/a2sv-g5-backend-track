# Task Management API Documentation

## Introduction

This API allows for the management of tasks, including creating, reading, updating, and deleting tasks. The data is stored in a MongoDB database.

## Prerequisites

Before running the API, ensure that you have the following installed:

- [Go](https://golang.org/doc/install)
- [MongoDB](https://www.mongodb.com/try/download/community) (either locally or via a cloud provider like [MongoDB Atlas](https://www.mongodb.com/cloud/atlas))
- [Gin Web Framework](https://github.com/gin-gonic/gin) (can be installed via `go get`)

## MongoDB Configuration

### 1. **Install MongoDB**

If you haven't installed MongoDB yet, follow the [installation guide](https://docs.mongodb.com/manual/installation/) for your operating system.

### 2. **Start MongoDB Server**

If you’re running MongoDB locally, start the MongoDB server by executing:

```sh
mongod
```

### 3. **Connecting to MongoDB**

The API connects to a MongoDB database. You can configure the connection string for MongoDB in your application code or environment variables.

- **Connection String Example**:
  
  ```sh
  mongodb://localhost:27017
  ```

- **Set Up MongoDB Database and Collection**:
  
  The application uses a database named `taskmanager` and a collection named `tasks`. These will be created automatically if they do not exist when the API is started.


## Setup and Running the API

### 1. **Clone the Repository**

```sh
git clone github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager.git
cd task_manager
```

### 2. **Install Dependencies**

```sh
go mod tidy
```

### 3. **Set Up MongoDB Connection**

Edit the `data/task_service.go` file or set the `MONGODB_URI` environment variable to specify your MongoDB connection string.

Example of connecting to MongoDB in `main.go`:

```go
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {log.Fatal(err)}
if client.Ping(context.TODO(), nil); err != nil {
  log.Fatal(err)
}
collection = client.Database("taskmanager").Collection("tasks")
```

### 4. **Run the API**

```sh
go run main.go
```

## API Endpoints

### Get Tasks

**Endpoint**: `GET /tasks`

This endpoint retrieves a list of tasks.

### Response

- **Status Code**: `200 OK`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  [
      {
          "id": "string",
          "title": "string",
          "description": "string",
          "due_date": "string",
          "status": "string"
      }
  ]
  ```

### Retrieve Task Details

**Endpoint**: `GET /tasks/:id`

This endpoint retrieves details of a specific task identified by its ID.

### Request

- **URL Parameters**: `id` (string) - The ID of the task to retrieve.

### Response

- **Status Code**: `200 OK`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "id": "string",
      "title": "string",
      "description": "string",
      "due_date": "string",
      "status": "string"
  }
  ```

If the task is not found:

- **Status Code**: `404 Not Found`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "error": "Task not found"
  }
  ```

### Create a Task

**Endpoint**: `POST /tasks`

This endpoint creates a new task.

### Request

- **Content Type**: `application/json`
- **Request Body**:

  ```json
  {
      "title": "string",
      "description": "string",
      "due_date": "string",
      "status": "string"
  }
  ```

### Response

- **Status Code**: `201 Created`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "message": "Task created successfully"
  }
  ```

### Update Task

**Endpoint**: `PUT /tasks/:id`

This endpoint updates a specific task by its ID.

### Request

- **Method**: `PUT`
- **URL Parameters**: `id` (string) - The ID of the task to update.
- **Content Type**: `application/json`
- **Request Body**:

  ```json
  {
      "title": "string",
      "description": "string",
      "due_date": "string",
      "status": "string"
  }
  ```

### Response

- **Status Code**: `200 OK`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "message": "Task updated successfully"
  }
  ```

### Delete Task

**Endpoint**: `DELETE /tasks/:id`

This endpoint deletes a specific task by its ID.

### Request

- **URL Parameters**: `id` (string) - The ID of the task to delete.

### Response

- **Status Code**: `200 OK`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "message": "Task deleted successfully"
  }
  ```
