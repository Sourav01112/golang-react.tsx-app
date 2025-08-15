# Todo App

A full-stack todo application built with Go (Fiber) backend and React (TypeScript) frontend.

## Features

- ✅ Create new todos
- ✅ Mark todos as complete
- ❌ Delete todos
- 📱 Responsive UI with Chakra UI
- 🔄 Real-time updates with React Query

## Tech Stack

**Backend:**
- Go with Fiber framework
- MongoDB for data storage
- CORS middleware for cross-origin requests

**Frontend:**
- React with TypeScript
- Chakra UI for styling
- React Query for state management
- React Icons

## API Endpoints

- `GET /health` - Health check
- `GET /api/get-todo` - Fetch all todos
- `POST /api/create-todo` - Create new todo
- `PATCH /api/update-todo/:id` - Mark todo as complete
- `DELETE /api/delete-todo/:id` - Delete todo

## Setup

1. **Backend Setup:**
   ```bash
   # Install dependencies
   go mod tidy
   
   # Set environment variables
   cp .env.example .env.local
   # Add your MONGODB_URI
   
   # Run server
   go run main.go
   ```

2. **Frontend Setup:**
   ```bash
   # Install dependencies
   npm install
   
   # Start development server
   npm run dev
   ```

## Environment Variables

- `MONGODB_URI` - MongoDB connection string
- `PORT` - Server port (default: 8000)
- `ENV` - Environment mode (development/production)

## Docker

The backend includes a multi-stage Dockerfile for containerized deployment:

```bash
# Build and run with Docker
docker build -t todo-app .
docker run -p 8080:8080 -e MONGODB_URI=your_uri todo-app
```

## Production

The app serves the React build files statically in production mode and includes health checks for container orchestration.