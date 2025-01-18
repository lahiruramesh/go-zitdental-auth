# Auth Application with GO

Authentication application using Zitadel OAuth2 service.


## Backend Setup

```bash
# Install dependencies
cd backend

go mod tidy

# Configure environment
cp .env.example .env

# Run server
go run main.go
```

## Frontend Setup

```bash
# Install dependencies
cd web
npm install

# Run development server
npm run dev
```


## API Endpoints

### Authentication Routes

| Method | Endpoint | Description | Request/Response |
|--------|----------|-------------|------------------|
| GET | `/api/auth/login` | Initiate login flow | Query: `?username=john.doe` |
| GET | `/api/oauth/callback` | OAuth callback handler | Query: `?code=xyz123` |
| GET | `/api/allowedUsers` | List allowed users | Response: `{ "results": [] }` |
| GET | `/api/profile` | Get user profile | Auth: `Bearer token` |

### Protected Routes

All routes require `Authorization: Bearer <token>` header

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| GET | `/api/profile` | Get authenticated user | None |