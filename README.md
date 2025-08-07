# Go Hexagonal Architecture

## Running Instructions

1. Make sure your Go version is 1.24 or higher

    ```sh
    go -version
    ```

2. Copy `.env.example` to `.env` and adjust accordingly

    ```sh
    cp .env.example .env
    ```

3. Install the required Go dependencies

    ```sh
    go mod download
    ```

4. Run the program

    ```sh
    ./run.sh
    ```

## Features

Auth:
- JWT
- OAuth2

RBAC:
- user
- admin

APIs:
- Store API (user)
- Payment gateway
- User CRUD (admin)
- Product CRUD (admin)

## Tech Stacks

API Framework:
- Gin

DB:
- Postgres
- GORM

Caching:
- Redis

Message Broker:
- RabbitMQ

CI / CD:
- Github Action

Libraries:
- github.com/gin-gonic/gin
- github.com/githubnemo/CompileDaemon
- github.com/joho/godotenv
- gorm.io/gorm
- gorm.io/driver/postgres
- github.com/google/uuid
