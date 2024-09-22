# Taurus App

Taurus App is a Go-based web service designed to manage financial expenses with a focus on supporting future mobile applications. This project is currently **confidential** and is in active development. It uses **MySQL** as the database backend and **Goose** for database migrations.

## Features

- **Go 1.23**: Built using the latest features of Go for high performance and scalability.
- **Docker**: The application is containerized for easy deployment and reproducibility.
- **Database Migrations**: Goose is used for handling schema migrations with MySQL as the database.
- **Minimal Docker Image**: Multi-stage Docker builds are used to keep the image size small for production.

## Prerequisites

- **Go 1.23** (For local development)
- **Docker** (To build and run the project in a container)
- **Goose** (For managing database migrations)
- **MySQL** (As the database for this project)

## Getting Started

### Clone the repository

```bash
git clone https://github.com/mrdhira/project-taurus.git
cd taurus-app
```

## Running the App with Docker

1. **Build the Docker Image**:

   ```
   docker build -t taurus-app .
   ```

2. **Run the Docker Container**:

   ```
   docker run -p 8000:8000 taurus-app
   ```

   The service will be available on `http://localhost:8000`.

## Running the App Locally

1. **Install Dependencies**:

   ```
   go mod download
   ```

2. **Run the Application**:

   ```
   go run ./cmd/serveHttp
   ```

   The application should now be running on `http://localhost:8000`.

## Database Migrations

Taurus App uses **Goose** to manage MySQL database migrations. Migrations are stored in the `db/migrations` directory.

### Setting Up Goose

To install Goose:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Creating a New Migration

To create a new migration file:

```
make migration-new
```

### Running Migrations

To apply the latest migrations:

```
make migration-up
```

To roll back the most recent migration:

```
make migration-down
```

To check the current migration status:

```
make migration-status
```

## Configuration

You can adjust the following settings as needed:

- **`DB_DSN`**: The MySQL database connection string can be modified in the `Makefile` for your environment.
- **`Ports`**: The default port exposed by the app is `8000`. You can adjust this in the Dockerfile or when running the Docker container.

## Building and Running the App

To build the Go binary for production:

```
go build -o taurus ./main.go
```

Run the binary:

```
./taurus serveHttp
```

## License

This project is licensed under the **Apache License, Version 2.0, January 2004**. See the [LICENSE](./LICENSE) file for more details.

## Confidentiality

Taurus App is a **confidential project** focusing on financial expense management and will support mobile applications in the future. Please do not share any details or code from this project without proper authorization.
