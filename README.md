# WhatsApp Clone Backend

A backend server implementation for a WhatsApp-like messaging application.

## Features

- Real-time messaging
- User authentication
- Message persistence
- Chat room management
- File sharing support

## Tech Stack

- Go
- gin-gonic/gin (HTTP framework)
- PostgreSQL (or other SQL database; configurable via internal/config)
- Redis (cache / pub-sub)
- Apache Kafka (event streaming)
- WebSockets (real-time messaging)
- JWT (authentication)
- Docker & Docker Compose

## Installation
### Prerequisites

- Go 1.20 or higher
- Docker and Docker Compose
- PostgreSQL
- Redis
- Apache Kafka

### Steps

```bash
# Clone the repository
git clone https://github.com/yourusername/whatsapp-clone-backend.git

# Navigate to project directory
cd whatsapp-clone-backend

# Install Go dependencies
go mod download

# Configure environment variables
cp .env.example .env

# Start infrastructure services
docker-compose up -d

# Run the server
go run cmd/server/main.go
```

For development:

```bash
# Run tests
go test ./...

# Build binary
go build -o whatsapp-server cmd/server/main.go
```

## API Endpoints

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - User login
- `GET /api/chats` - Get user chats
- `POST /api/messages` - Send message


## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT
