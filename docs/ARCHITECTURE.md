# Project Code Organization Guide (Scalable Architecture)

This document explains the purpose of the key directories in our architecture: `internal/pkg`, `internal/core`, and `internal/modules`. It includes examples for future growth to ensure the project remains maintainable and scalable.

## 1. `internal/pkg` (Generic Utilities)
**Purpose**: Contains universally reusable code that is **domain-agnostic**. These packages act as your internal open-source library. They should never import from `core` or `modules`.

**Rules**:
- No dependencies on business logic.
- Pure functions or wrappers around external libraries.

**Current Examples**:
- **`hashing`**: Password hashing.
- **`logger`**: Structured logging.

**🚀 Future Growth Examples**:
- **`mailer`**: A wrapper for sending emails (SMTP, SendGrid, SES) without knowing *what* email is being sent.
- **`storage`**: specific interface to upload files to AWS S3, Google Cloud Storage, or MinIO.
- **`pdf`**: Utility to generate PDF invoices from HTML templates.
- **`excel`**: Utility to parse or generate Excel reports.
- **`queue`**: A generic wrapper for RabbitMQ or Kafka.

---

## 2. `internal/core` (Application Infrastructure)
**Purpose**: Contains the essential "glue" code and infrastructure specific to **this application**. It sets up the technical foundation.

**Rules**:
- Can depend on configuration (`env`).
- Sets up global dependencies.
- Acts as the framework layer.

**Current Examples**:
- **`config`**: Loads `.env`.
- **`database`**: Connects to Postgres.
- **`server`**: HTTP server setup.

**🚀 Future Growth Examples**:
- **`cache`**: Redis client setup and specific caching strategies (e.g., "Look-aside" cache helpers).
- **`worker`**: Setup for background job workers (e.g., Asynq, Machinery) to process tasks asynchronously.
- **`sentry`**: Error tracking setup and global panic recovery.
- **`metrics`**: Prometheus exporter setup for monitoring API performance.

---

## 3. `internal/modules` (Business Logic)
**Purpose**: Contains the actual **features** of your application. Organized by Domain (DDD-lite). This is where 90% of your daily work happens.

**Rules**:
- Highly cohesive. Everything related to a feature stays together.
- Communicate via Services or Interfaces (loose coupling).

**Structure within a Module**:
```text
internal/modules/payment/
├── controller/       # HTTP Handlers (POST /pay)
├── service/          # Business Logic (Calculate tax, Call Stripe)
├── repository/       # DB Access (Save transaction record)
├── domain/           #/ Structs (Payment, Transaction)
└── dto/              # Request/Response structs (PayRequest, PayResponse)
```

**Current Examples**:
- **`auth`**, **`user`**, **`iam`**.

**🚀 Future Growth Examples**:
- **`notification`**: Listens for events (UserRegistered) and uses `core/worker` + `pkg/mailer` to send emails.
- **`payment`**: Handles payment gateways (Stripe/PayPal), transaction history, and currency conversion.
- **`order`**: Manages cart, checkout flows, and order status state machines.
- **`inventory`**: Manages stock levels, reservations, and warehouse logic.
- **`report`**: Aggregates data from other modules to generate monthly analytics.

---

## Why this Structure Scales?

1.  **Maintainability**: If you need to change how you send emails, you only touch `pkg/mailer`. If you need to change how Login works, you only touch `modules/auth`.
2.  **Testability**: Since modules are separated, you can easily write unit tests for `service` logic by mocking `repository` and `pkg`.
3.  **Onboarding**: A new developer working on "Payments" only needs to look at `internal/modules/payment`.
4.  **Microservices Ready**: If the `payment` module becomes too complex or needs to scale independently, you can lift the entire folder out into its own repository/microservice easily because it's already isolated.
