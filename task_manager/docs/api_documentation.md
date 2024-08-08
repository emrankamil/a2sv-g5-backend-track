
# Task Management API Documentation

## Get Tasks

**Endpoint**: `GET /tasks`

This endpoint is used to retrieve a list of tasks.

### Request

- **Query Parameters**: No query parameters required.

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

Each task object includes the following attributes:

- `id` (string): The unique identifier for the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `due_date` (string): The due date of the task.
- `status` (string): The status of the task.

---

## Retrieve Task Details

**Endpoint**: `GET /tasks/:id`

This endpoint is used to retrieve details of a specific task identified by its ID.

### Request

- **URL Parameters**: `id` (string) - The ID of the task to retrieve.
- **Request Body**: This is a `GET` request and does not require a request body.

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

In case the requested task is not found, the response will have:

- **Status Code**: `404 Not Found`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "error": "string"
  }
  ```

---

## Create a Task

**Endpoint**: `POST /tasks`

This endpoint is used to create a new task.

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

The request body requires the following parameters:

- `title` (string, required): The title of the task.
- `description` (string, required): The description of the task.
- `due_date` (string, required): The due date of the task.
- `status` (string, required): The status of the task.

### Response

- **Status Code**: `201 Created`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "message": "Task created successfully"
  }
  ```

---

## Update Task

**Endpoint**: `PUT /tasks/:id`

This endpoint is used to update a specific task by providing the task ID in the URL.

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

The request requires a JSON object in the body with the following parameters:

- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `due_date` (string): The due date of the task.
- `status` (string): The status of the task.

### Response

- **Status Code**: `200 OK`
- **Content Type**: `application/json`
- **Response Body**:

  ```json
  {
      "message": "Task updated successfully"
  }
  ```

---

## Delete Task

**Endpoint**: `DELETE /tasks/:id`

This endpoint is used to delete a specific task identified by its ID.

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

---

This Markdown documentation should now be more clear and readable for developers to use with your Task Management API.

## Running the API

### Run the server:
```sh
go run main.go
```

### Use Postman to test each endpoint:

- **GET** `/tasks`
- **GET** `/tasks/:id`
- **POST** `/tasks` with JSON body:
  ```json
  {
    "title": "Task 1",
    "description": "Description 1",
    "due_date": "2024-12-31",
    "status": "pending"
  }
  ```
- **PUT** `/tasks/:id` with JSON body:
  ```json
  {
    "title": "Updated Task",
    "description": "Updated Description",
    "due_date": "2024-12-31",
    "status": "completed"
  }
  ```
- **DELETE** `/tasks/:id`