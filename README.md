# Risks API

## Prerequisites (dependencies)

- Docker ^23.0
- Docker Compose plugin ^2.17

## How to run

Run:
```
docker compose up --build
```

Then navigate to http://localhost:8080/swagger/index.html to see the Swagger UI and to try endpoints. Use the "Try it out" button.

## Running tests

Run:
```
docker compose run --build --rm risks-api go test -v ./...
```
