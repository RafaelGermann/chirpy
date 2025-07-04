# Chirpy - A Twitter-like API

A RESTful API built with Go that mimics Twitter's functionality, created as a learning project from Boot.dev. This project demonstrates modern Go web development practices including authentication, database operations, and API design.

## Features

- **User Management**: Registration, authentication, and profile updates
- **Chirps**: Create, read, and delete posts
- **JWT Authentication**: Secure token-based authentication with refresh tokens
- **Content Moderation**: Automatic filtering of inappropriate words
- **Webhook Integration**: External service integration for user upgrades
- **Metrics**: Admin dashboard for monitoring API usage
- **PostgreSQL Database**: Persistent data storage with SQLC for type-safe queries

## Tech Stack

- **Language**: Go 1.24.4
- **Database**: PostgreSQL
- **ORM**: SQLC (SQL Compiler)
- **Authentication**: JWT tokens with refresh tokens
- **Password Hashing**: bcrypt
- **Environment**: godotenv for configuration

## API Endpoints

### Authentication Endpoints

#### `POST /api/users`
Create a new user account.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response**: User object with ID, email, and timestamps
- **Status**: 201 Created

#### `POST /api/login`
Authenticate a user and receive access tokens.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
- **Response**: User object with JWT token and refresh token
- **Status**: 200 OK

#### `POST /api/refresh`
Refresh an expired JWT token using a refresh token.
- **Headers**: `Authorization: Bearer <refresh_token>`
- **Response**: New JWT token
- **Status**: 200 OK

#### `POST /api/revoke`
Revoke a refresh token (logout).
- **Headers**: `Authorization: Bearer <refresh_token>`
- **Status**: 204 No Content

### User Management

#### `PUT /api/users`
Update user profile information.
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Request Body**:
  ```json
  {
    "email": "newemail@example.com",
    "password": "newpassword123"
  }
  ```
- **Response**: Updated user object
- **Status**: 200 OK

### Chirps (Posts)

#### `GET /api/chirps`
Retrieve all chirps with optional filtering and sorting.
- **Query Parameters**:
  - `author_id`: Filter chirps by user ID
  - `sort`: Sort order (`asc` or `desc`, default: `asc`)
- **Response**: Array of chirp objects
- **Status**: 200 OK

#### `GET /api/chirps/{chirpID}`
Retrieve a specific chirp by ID.
- **Path Parameters**: `chirpID` - UUID of the chirp
- **Response**: Chirp object
- **Status**: 200 OK

#### `POST /api/chirps`
Create a new chirp (requires authentication).
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Request Body**:
  ```json
  {
    "body": "Hello, world! This is my first chirp!"
  }
  ```
- **Response**: Created chirp object
- **Status**: 201 Created
- **Features**:
  - Maximum 140 characters
  - Automatic content filtering (replaces inappropriate words with "****")

#### `DELETE /api/chirps/{chirpID}`
Delete a chirp (only by the author).
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Path Parameters**: `chirpID` - UUID of the chirp
- **Status**: 204 No Content

### Webhooks

#### `POST /api/polka/webhooks`
Handle external webhook events (e.g., user upgrades).
- **Headers**: `Authorization: ApiKey <api_key>`
- **Request Body**:
  ```json
  {
    "event": "user.upgraded",
    "data": {
      "user_id": "uuid-here"
    }
  }
  ```
- **Status**: 204 No Content

### Admin Endpoints

#### `GET /admin/metrics`
Admin dashboard showing API usage statistics.
- **Response**: HTML page with hit counter
- **Status**: 200 OK

#### `POST /admin/reset`
Reset the application (development only).
- **Status**: 200 OK (dev environment only)

### Health Check

#### `GET /api/healthz`
Health check endpoint.
- **Response**: "OK"
- **Status**: 200 OK

## Data Models

### User
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "is_chirpy_red": false
}
```

### Chirp
```json
{
  "id": "uuid",
  "body": "Chirp content here",
  "user_id": "uuid",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Environment Variables

Create a `.env` file with the following variables:

```env
DB_URL=postgres://username:password@localhost:5432/chirpy
PLATFORM=dev
JWT_SECRET=your-secret-key-here
POLKA_KEY=your-webhook-api-key
```

## Setup Instructions

1. **Install Dependencies**:
   ```bash
   go mod download
   ```

2. **Database Setup**:
   - Install PostgreSQL
   - Create a database named `chirpy`
   - Run the SQL migrations in the `sql/schema/` directory

3. **Environment Configuration**:
   - Copy the environment variables above to a `.env` file
   - Update the values according to your setup

4. **Run the Application**:
   ```bash
   go run .
   ```

5. **Access the API**:
   - The server will start on port 8080
   - API base URL: `http://localhost:8080/api/`
   - Admin dashboard: `http://localhost:8080/admin/metrics`

## Security Features

- **Password Hashing**: All passwords are hashed using bcrypt
- **JWT Authentication**: Secure token-based authentication
- **Refresh Tokens**: Long-lived refresh tokens for better UX
- **Content Filtering**: Automatic filtering of inappropriate content
- **Authorization**: Users can only delete their own chirps
- **API Key Protection**: Webhook endpoints protected with API keys

## Development Notes

- Built as a learning project for Go web development

## License

This project was created as part of the Boot.dev Go course for educational purposes. 