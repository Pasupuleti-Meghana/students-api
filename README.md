StudentHub – Student Management REST API (Go)

A simple backend service built with Golang and SQLite that provides REST APIs to create and retrieve student records. This project demonstrates backend fundamentals such as HTTP handling, request validation, structured logging, and database interaction.

Features

Create a new student

Get a student by ID

Retrieve all students

Request validation

Structured logging

SQLite database integration

Clean project structure

Tech Stack

Go (net/http)

SQLite

JSON

go-playground/validator

Structured logging (slog)

Project Structure
students-api
│
├── cmd/students-api
│   └── main.go
│
├── internal
│   ├── http
│   │   └── handlers
│   │       └── student
│   │
│   ├── storage
│   │
│   ├── types
│   │
│   └── utils
│
├── config
│   └── local.yaml
│
└── storage
    └── storage.db
API Endpoints
Create Student

POST /students

Request

{
  "name": "Meghana",
  "email": "meghana@gmail.com",
  "age": 22
}

Response

{
  "id": 1
}
Get Student by ID

GET /students/{id}

Example

GET /students/1

Response

{
  "id": 1,
  "name": "Meghana",
  "email": "meghana@gmail.com",
  "age": 22
}
Get All Students

GET /students

Response

[
  {
    "id": 1,
    "name": "Meghana",
    "email": "meghana@gmail.com",
    "age": 22
  }
]

How to Run the Project
1. Clone the repository
git clone https://github.com/yourusername/students-api.git
2. Go to the project folder
cd students-api
3. Install dependencies
go mod tidy
4. Run the server
go run cmd/students-api/main.go -config config/local.yaml

Server will start on

http://localhost:8082
Example cURL Requests

Create Student

curl -X POST http://localhost:8082/students \
-H "Content-Type: application/json" \
-d '{"name":"Meghana","email":"meghana@gmail.com","age":22}'

Get Student

curl http://localhost:8082/students/1

Get All Students

curl http://localhost:8082/students
What I Learned

Building REST APIs in Go

Handling HTTP requests and responses

Request validation

Structuring backend projects

Working with SQLite databases

Debugging and error handling

Future Improvements

Update student API

Delete student API

Pagination

Authentication (JWT)

Docker support

Unit tests

Author
Meghana Pasupuleti