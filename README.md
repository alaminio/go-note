# Go Note & Todo API

A simple REST API for managing notes and todos built with Go, Gin, and GORM.

## Project Structure

```
project/
│
├── cmd/
│   └── server/
│       └── main.go          # Entry point: wires DB + router
│
├── internal/                # Private app code (not importable outside module)
│   ├── api/                 # HTTP layer (Gin handlers + routes)
│   │   ├── router.go
│   │   ├── note_handler.go
│   │   └── todo_handler.go
│   │
│   ├── note/                # Domain logic for "notes"
│   │   ├── model.go         # GORM model (Note struct)
│   │   ├── repository.go    # DB access (CRUD with GORM)
│   │   └── service.go       # Business logic
│   │
│   ├── todo/                # Domain logic for "todos"
│   │   ├── model.go         # GORM model (Todo struct)
│   │   ├── repository.go    # DB access (CRUD with GORM)
│   │   └── service.go       # Business logic
│   │
│   └── db/
│       └── db.go            # ConnectDB, AutoMigrate
│
├── pkg/                     # Optional reusable utilities
│   └── logger/              # Example: custom logging
│
├── migrations/              # SQL migration files (if you use them)
│
├── configs/                 # Config files (env, yaml, json)
│
├── go.mod
├── go.sum
└── README.md
```

## Dependencies

- Gin: HTTP web framework
- GORM: ORM library for Go
- SQLite: Database (can be changed to PostgreSQL, MySQL, etc.)

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

```bash
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Notes
- `POST /api/notes` - Create a new note
- `GET /api/notes` - Get all notes
- `GET /api/notes/:id` - Get a specific note
- `PUT /api/notes/:id` - Update a note
- `DELETE /api/notes/:id` - Delete a note

### Todos
- `POST /api/todos` - Create a new todo
- `GET /api/todos` - Get all todos
- `GET /api/todos/:id` - Get a specific todo
- `PUT /api/todos/:id` - Update a todo
- `DELETE /api/todos/:id` - Delete a todo
- `PATCH /api/todos/:id/toggle` - Toggle todo completion status

## Example Usage

### Notes
Create a note:
```bash
curl -X POST http://localhost:8080/api/notes \
  -H "Content-Type: application/json" \
  -d '{"title": "My Note", "content": "This is my first note"}'
```

Get all notes:
```bash
curl http://localhost:8080/api/notes
```

### Todos
Create a todo:
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Buy groceries", "description": "Milk, bread, eggs"}'
```

Get all todos:
```bash
curl http://localhost:8080/api/todos
```

Toggle todo completion:
```bash
curl -X PATCH http://localhost:8080/api/todos/1/toggle