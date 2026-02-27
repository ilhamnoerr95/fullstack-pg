# 🚀 Full Stack Monorepo (Turborepo + Golang + Vue 3)

## 📦 Overview

1. Project ini adalah Full Stack Monorepo menggunakan:

2. Turborepo sebagai monorepo task runner

3. Backend: Golang (DDD Architecture)

4. Frontend: Vue 3 (Composition API + Atomic Design)

## 🏗️ Architecture Full Stack

```code
root/
├── apps/
│   ├── backend/   → Golang (DDD)
│   └── frontend/  → Vue 3
├── .env
├── docker-compose.yml
├── turbo.json
├── package.json
└── README.md
```

## 🧠 Backend Stack

Backend dibangun menggunakan:

- Golang
- DDD (Domain Driven Design)
- Gin (HTTP Framework)
- Swagger (API Docs)
- JWT v5 (Authentication)
- Redis (Caching)
- Clean Architecture separation:
- domain
- repository
- service
- handler
- infrastructure

## 🎨 Frontend Stack

Frontend dibangun menggunakan:

- Vue 3 (Composition API)
- Atomic Design Pattern
- Pinia (State Management)
- Vue Router
- Ant Design Vue
- Day.js

## ⚙️ Requirements

Pastikan sudah terinstall:

1. Node.js ≥= 20
2. npm
3. Go ≥ 1.22
4. Docker
5. Docker Compose

## 🛠️ Installation

### 1️⃣ Clone Repository

```bash
git clone git@github.com:ilhamnoerr95/fullstack-pg.git
cd your-repo
```

### 2️⃣ Install Dependencies (Root)

```bash
npm install
```

### 3️⃣ Setup Backend

```bash
cd apps/backend
go mod tidy
```

### 4️⃣ Setup Frontend

```bash
cd apps/frontend
npm install
```

### 5️⃣ Setup Redis (Docker)

Di root project:

```bash
docker compose up -d
```

cek

```bash
docker ps
```

Redis akan berjalan di:

```bash
localhost:6379
```

## 🚀 Running the Project

### 🔥 Run All Apps (Monorepo Mode)

Dari root:

```bash
npm run dev
```

### 🧩 Run Backend Only

```bash
cd apps/backend
npm run dev
```

atau

```bash
go run main.go
```

Swagger:

```code
http://localhost:8080/docs
```

### 🎨 Run Frontend Only

```bash
cd apps/frontend
npm run dev
```

Frontend berjalan di:

```code
http://localhost:5173
```

## 🔐 Authentication

Backend menggunakan JWT v5.
Protected endpoint harus menyertakan:

```code
Authorization: Bearer <token>
```

## 🗂️ Backend Architecture (DDD)

```bash
apps/backend/
├── main.go                         # Entry point
│
├── docs/                            # Swagger HTML
│   └── swagger.html
│
├── internal/
│   │
│   ├── domain/                      # Pure business entities
│   │   ├── user.go
│   │   ├── payment.go
│   │
│   ├── repository/                  # Interfaces (Contracts)
│   │   ├── user_repository.go
│   │   └── payment_repository.go
│   │
│   ├── service/                     # Business logic (Usecases)
│   │   ├── auth_service.go
│   │   ├── payment_service.go
│   │
│   ├── handler/                     # HTTP layer (Gin)
│   │   ├── auth_handler.go
│   │   ├── payment_handler.go
│   │
│   ├── middleware/                  # JWT, Auth, Logging
│   │   └── auth_middleware.go
│   │
│   ├── infrastructure/              # External implementations
│   │   ├── redis_client.go
│   │
│   └── lib/                       # Helpers (optional)
│       └── jwt.go
│
├── openapi.yaml                     # Swagger spec
├── .air.toml                        # Dev hot reload
├── go.mod
├── go.sum
└── Makefile
```

## 🔄 Backend Flow

```bash
HTTP Request
   ↓
Handler (Gin)
   ↓
Service (Business Logic)
   ↓
Repository Interface
   ↓
Infrastructure (DB / Redis)
```

## 🧱 Responsibility per Layer

```
| Layer          | Responsibility        |
| -------------- | --------------------- |
| domain         | Entity & core rules   |
| repository     | Interface abstraction |
| service        | Business logic        |
| handler        | HTTP + JSON           |
| middleware     | JWT protection        |
| infrastructure | Redis / DB impl       |

```

## ⚡ Redis Usage

Redis digunakan untuk:

- Caching payments
- Performance optimization
- TTL cache expiration

Flow:

```code
Check Redis
 ↓
Hit → return
Miss → query DB → save to Redis
```

## 🎨 Frontend Architecture (Atomic Design)

```code
components/
├── atoms/
├── molecules/
├── organisms/
├── templates/
```

## 🧑‍💻 Development Notes

- Follow DDD separation strictly
- Service layer should not depend on HTTP
- Redis logic only inside service/infrastructure
- Frontend components follow atomic hierarchy
- Avoid business logic inside Vue components

## 🏁 Summary

This monorepo combines:

- ⚡ High-performance Go backend
- 🎨 Scalable Vue 3 frontend
- 🧱 Clean DDD architecture
- 🚀 Turbo-powered monorepo workflow
- 🔥 Redis caching optimization
