### Task Management API Documentation

## Endpoints

### Get all tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Success Response:**
  - **Code:** 200
  - **Content:** 
    ```json
    [
      {
        "id": "1",
        "title": "Task 1",
        "description": "Description 1",
        "due_date": "2024-12-31",
        "status": "pending"
      }
    ]
    ```

### Get task by ID
- **URL:** `/tasks/:id`
- **Method:** `GET`
- **URL Params:** `id=[string]`
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "id": "1",
      "title": "Task 1",
      "description": "Description 1",
      "due_date": "2024-12-31",
      "status": "pending"
    }
    ```
- **Error Response:**
  - **Code:** 404
  - **Content:** 
    ```json
    {
      "error": "Task not found"
    }
    ```

### Create new task
- **URL:** `/tasks`
- **Method:** `POST`
- **Data Params:**
  ```json
  {
    "title": "Task 1",
    "description": "Description 1",
    "due_date": "2024-12-31",
    "status": "pending"
  }
  ```
- **Success Response:**
  - **Code:** 201
  - **Content:** 
    ```json
    {
      "id": "1",
      "title": "Task 1",
      "description": "Description 1",
      "due_date": "2024-12-31",
      "status": "pending"
    }
    ```

### Update task by ID
- **URL:** `/tasks/:id`
- **Method:** `PUT`
- **URL Params:** `id=[integer]`
- **Data Params:**
  ```json
  {
    "id":"2",
    "title": "Updated Task",
    "description": "Updated Description",
    "due_date": "2024-12-31",
    "status": "completed"
  }
  ```
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "id": "1",
      "title": "Updated Task",
      "description": "Updated Description",
      "due_date": "2024-12-31",
      "status": "completed"
    }
    ```
- **Error Response:**
  - **Code:** 404
  - **Content:**
    ```json
    {
      "error": "Task not found"
    }
    ```

### Delete task by ID
- **URL:** `/tasks/:id`
- **Method:** `DELETE`
- **URL Params:** `id=[integer]`
- **Success Response:**
  - **Code:** 200
  - **Content:**
    ```json
    {
      "message": "Task deleted"
    }
    ```
- **Error Response:**
  - **Code:** 404
  - **Content:**
    ```json
    {
      "error": "Task not found"
    }
    ```

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