<div align="center">

# ⚙️ Hospital API — Go Backend

[![Go](https://img.shields.io/badge/Go-1.22-blue?logo=go)](https://go.dev)
[![Fiber](https://img.shields.io/badge/Fiber-v2-blue)](https://gofiber.io)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![CI](https://github.com/zuhudo/hospital-app-backend/actions/workflows/go-ci.yml/badge.svg)](https://github.com/zuhudo/hospital-app-backend/actions/workflows/go-ci.yml)

High-performance REST API backend for the Hospital & Patient Management System.

[📱 Mobile App](https://github.com/zuhudo/hospital-app-mobile) • [🌐 Website](https://github.com/zuhudo/hospital-app-web) • [📊 Dashboard](https://github.com/zuhudo/hospital-app-dashboard) • [📖 Wiki](https://github.com/zuhudo/hospital-app-backend/wiki)

</div>

## 📋 Features

- 🔑 **JWT Authentication** — Secure login & registration
- 👥 **Patient CRUD** — Full patient management
- 👨‍⚕️ **Doctor Management** — Doctor profiles & schedules
- 📅 **Appointment Scheduling** — Book, cancel, list appointments
- 📋 **Medical Records** — Patient history & prescriptions
- 🛡️ **Middleware** — CORS, logging, recovery, request ID
- 🐳 **Docker** — Containerized deployment
- 📊 **Health Check** — `/health` endpoint

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+

### Installation

```bash
# Clone the repo
git clone https://github.com/zuhudo/hospital-app-backend.git
cd hospital-app-backend

# Copy environment file
cp .env.example .env

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

The API will be available at `http://localhost:3000`

### Build

```bash
# Build binary
go build -o server main.go

# Run binary
./server
```

### Docker

```bash
# Build image
docker build -t hospital-backend .

# Run container
docker run -p 3000:3000 hospital-backend
```

## 📡 API Endpoints

### Auth
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/auth/login` | User login |
| POST | `/api/auth/register` | User registration |

### Patients
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/patients` | List all patients |
| POST | `/api/patients` | Create patient |
| GET | `/api/patients/:id` | Get patient by ID |
| PUT | `/api/patients/:id` | Update patient |
| DELETE | `/api/patients/:id` | Delete patient |

### Doctors
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/doctors` | List all doctors |
| POST | `/api/doctors` | Create doctor |
| GET | `/api/doctors/:id` | Get doctor by ID |

### Appointments
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/appointments` | List appointments |
| POST | `/api/appointments` | Book appointment |
| GET | `/api/appointments/:id` | Get appointment |
| PUT | `/api/appointments/:id/cancel` | Cancel appointment |

### Medical Records
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/records/:patientId` | Get patient records |
| POST | `/api/records` | Create record |

## 📁 Project Structure

```
backend/
├── main.go                    # Entry point
├── cmd/server/                # Server command
├── internal/
│   ├── config/                # Configuration
│   ├── database/              # Database connection
│   ├── middleware/             # Auth, CORS, Logger
│   ├── models/                # Data models
│   ├── handlers/              # Request handlers
│   ├── routes/                # Route registration
│   └── utils/                 # Helpers
├── .env.example               # Environment template
└── Dockerfile                 # Docker config
```

## 🛠️ Tech Stack

- **Go** 1.22 — Programming language
- **Fiber v2** — Web framework
- **JWT** — Authentication
- **bcrypt** — Password hashing
- **Docker** — Containerization

## 📖 Documentation

- [Getting Started](https://github.com/zuhudo/hospital-app-backend/wiki/Getting-Started)
- [API Reference](https://github.com/zuhudo/hospital-app-backend/wiki/API-Reference)
- [Authentication](https://github.com/zuhudo/hospital-app-backend/wiki/Authentication)
- [Deployment](https://github.com/zuhudo/hospital-app-backend/wiki/Deployment)

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📄 License

MIT License — see [LICENSE](LICENSE) for details.
