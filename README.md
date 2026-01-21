# Modular Monolith Go Setup

A production-ready, modular monolith architecture for building scalable Go applications. This project serves as a robust foundation featuring clean architecture, essential infrastructure, and best practices.

## 🚀 Tech Stack

- **Language**: [Go 1.25+](https://go.dev/)
- **Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin)
- **Database**: [PostgreSQL](https://www.postgresql.org/) (via [GORM](https://gorm.io/))
- **Cache**: [Redis](https://redis.io/)
- **Migrations**: [Atlas](https://atlasgo.io/)
- **Documentation**: [Swagger](https://swagger.io/) (via [swaggo](https://github.com/swaggo/swag))
- **Live Reload**: [Air](https://github.com/air-verse/air)
- **Containerization**: [Docker](https://www.docker.com/)

## 🏗 Architecture

We follow a **Modular Monolith** approach, splitting the application into distinct feature modules (e.g., `user`, `auth`, `iam`) while sharing a robust core infrastructure.

For a detailed deep-dive into our code organization (`pkg` vs `core` vs `modules`), please read:
👉 **[Architecture Guide](docs/ARCHITECTURE.md)**

## 🛠 Getting Started

### Prerequisites
- Go 1.25+
- Docker & Docker Compose
- Make

### 1. Installation
```bash
git clone https://github.com/your-repo/go-setup.git
cd go-setup
go mod download
```

### 2. Environment Setup
Copy the example environment file and configure your secrets:
```bash
cp .env.example .env
```
> **Security Note**: Ensure you set strong values for `JWT_SECRET`, `REFRESH_TOKEN_SECRET`, and `PASSWORD_HASH_SECRET`.

### 3. Start Infrastructure
Launch Postgres and Redis containers:
```bash
make docker-infra-up
```

### 4. Database Migration
Apply the database schema:
```bash
make migrate
```

### 5. Seed Data (Optional)
Populate the database with initial demo data (Roles, Permissions, default Admin/User):
```bash
make seed
```
> **Default Accounts**:
> - Admin: `admin@example.com` / `admin@1901`
> - User: `user@example.com` / `user@1901`

### 6. Run Application
Start the server with live reload (Air):
```bash
make dev
```
The API will be available at `http://localhost:8080`.

## 📚 API Documentation

Once the application is running, you can access the interactive Swagger documentation at:
**[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

To update the documentation after changing code comments:
```bash
make swagger
```

## ⚡ Command Reference

| Command | Description |
| :--- | :--- |
| `make dev` | Run the app with live reload (Air) |
| `make run` | Run the app directly (`go run`) |
| `make build` | Build the binary to `bin/app` |
| `make docker-infra-up` | Start Postgres & Redis containers |
| `make docker-infra-down` | Stop and remove infrastructure containers |
| `make seed` | **Run Database Seeders** (Users, Roles, Credentials) |
| `make migrate` | Apply database migrations |
| `make migrate-generate` | Generate migration files based on GORM models |
| `make swagger` | Generate Swagger documentation |
