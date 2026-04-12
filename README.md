# Go Opportunities

Go Opportunities is a simple REST API for managing job openings. It is built with **Go**, **Gin**, **GORM**, and **SQLite**, and includes **Swagger UI** for interactive API documentation.

## Overview

This project exposes a small CRUD API for job opportunities, allowing you to:

1. create a job opening
2. list all openings
3. fetch a single opening by ID
4. update an existing opening
5. delete an opening

The application stores data in a local SQLite database at `db/main.db`, which is created automatically on first run.

## Tech Stack

| Layer | Technology |
| --- | --- |
| Language | Go |
| HTTP framework | Gin |
| ORM | GORM |
| Database | SQLite |
| API docs | Swagger / swaggo |

## Project Structure

```text
.
├── config/    # database and logger setup
├── db/        # SQLite database file
├── docs/      # generated Swagger documentation
├── handler/   # HTTP handlers, requests, and responses
├── router/    # route registration and server bootstrap
├── schemas/   # database models and response schemas
└── main.go    # application entry point
```

## Getting Started

### Prerequisites

- Go installed
- A Go version compatible with the one declared in `go.mod`

### Installation

```bash
git clone https://github.com/ArctisDev/go-opportunities.git
cd go-opportunities
go mod download
```

### Run the API

```bash
go run main.go
```

The server starts on:

```text
http://localhost:8080
```

## API Base URL

```text
http://localhost:8080/api/v1
```

## Available Endpoints

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | `/api/v1/openings` | List all job openings |
| GET | `/api/v1/opening?id={id}` | Get a single opening |
| POST | `/api/v1/opening` | Create a new opening |
| PUT | `/api/v1/opening?id={id}` | Update an opening |
| DELETE | `/api/v1/opening?id={id}` | Delete an opening |

## Opening Payload

Example request body used for creating an opening:

```json
{
  "role": "Backend Go Developer",
  "company": "Acme Inc.",
  "description": "Build and maintain backend services in Go.",
  "remote": true,
  "location": "Sao Paulo, Brazil",
  "salary": 120000,
  "url": "https://example.com/jobs/backend-go-developer"
}
```

For updates, send only the fields you want to change in the request body and provide the target opening ID in the query string.

## Swagger Documentation

Swagger UI is available at:

```text
http://localhost:8080/swagger/index.html
```

## Example Requests

### Create an opening

```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Backend Go Developer",
    "company": "Acme Inc.",
    "description": "Build and maintain backend services in Go.",
    "remote": true,
    "location": "Sao Paulo, Brazil",
    "salary": 120000,
    "url": "https://example.com/jobs/backend-go-developer"
  }'
```

### List all openings

```bash
curl http://localhost:8080/api/v1/openings
```

### Get one opening

```bash
curl "http://localhost:8080/api/v1/opening?id=1"
```

### Update one opening

```bash
curl -X PUT "http://localhost:8080/api/v1/opening?id=1" \
  -H "Content-Type: application/json" \
  -d '{
    "salary": 130000,
    "remote": false
  }'
```

### Delete one opening

```bash
curl -X DELETE "http://localhost:8080/api/v1/opening?id=1"
```

## Notes

- The database schema is auto-migrated when the application starts.
- The SQLite file is stored locally, which makes the project convenient for learning, prototyping, and small internal tools.