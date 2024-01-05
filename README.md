# Golang Review API

This is the API for managing reviews and comments. It allows users to create reviews, view existing reviews, and post comments on reviews.

## Getting Started

### Prerequisites

Before running the application, make sure you have the following installed:

- Go (Golang)
- MongoDB
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/golang-review-api.git
   ```

2. Change into the project directory:

   ```bash
   cd golang-review-api
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

### Configuration

Make sure to set up your MongoDB connection string in the `.env` file or as an environment variable:

```env
MONGO_URI=mongodb://your-username:your-password@localhost:27017/golang_backend_review
SECRET_KEY=your-secret-key
```

### Running the Application

Run the following command to start the API server:

```bash
go run main.go
```

The API server will be accessible at [http://localhost:8080](http://localhost:8080).

### API Documentation

Swagger documentation is available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html). You can use the Swagger UI to explore and test the API endpoints.

## API Endpoints

### Create a Review

**Endpoint:** `POST /reviews`

Create a new review with a title and description.

### Get All Reviews

**Endpoint:** `GET /reviews`

Get a list of all reviews.

### Get Review by ID

**Endpoint:** `GET /reviews/{id}`

Get details of a specific review by providing its ID.

### Post Comment

**Endpoint:** `POST /reviews/{id}/comment`

Post a comment on a specific review by providing its ID.

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. Include the JWT token in the `Authorization` header for authenticated routes.
