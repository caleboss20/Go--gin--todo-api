# Go Todo API
A RESTful Todo API built with Golang for learning backend development and API design.

## Project Structure
```bash
.
├── cmd
├── internal
│   ├── handlers
│   └── models
```
## Features
- Create Todos
- Retrieve Todos
- Update Todos
- Delete Todos

## Tech Stack
- Golang
- net/http
- gin
 
## Getting Started

### Clone the Repository
```bash
git clone https://github.com/caleboss20/go--gin--todo-api.git
```
### Navigate to the Project Directory
```bash
cd go-todo-api
```
### Run the Application
```bash
go run ./cmd
```
## API Endpoints

| Method | Endpoint      | Description       |
| ------- | ------------- | ----------------- |
| GET     | `/todos`      | Retrieve all todos |
| POST    | `/todos`      | Create a new todo |
| PUT     | `/todos/:id`  | Update a todo     |
| DELETE  | `/todos/:id`  | Delete a todo     |

## Author
Developed as part of a backend engineering learning journey using Go.
