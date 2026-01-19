# Smart Conversation Backend

Go backend for smart conversation application with GLM-4.6v API integration.

## Features

- Text-based conversation with GLM-4.6v model
- Image upload and analysis
- Multi-modal conversation support

## API Endpoints

- POST `/api/chat` - Send conversation messages with optional image
- POST `/api/upload` - Upload image for processing
- GET `/*` - Serve frontend static files

## Setup

1. Install dependencies:
```bash
cd backend
go mod tidy
```

2. Run the server:
```bash
go run main.go
```

The server will start on port 8080.