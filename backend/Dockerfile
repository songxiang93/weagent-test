FROM golang:1.21-alpine

# Install git for go modules
RUN apk add --no-cache git

WORKDIR /app

# Copy all files first
COPY . .

# Generate go.sum and tidy
RUN go mod download && go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

EXPOSE 8080

CMD ["./server"]