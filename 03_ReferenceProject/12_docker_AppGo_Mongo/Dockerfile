# Start from the official golang image, version 1.17
FROM golang:1.17

# Set the working directory to /app
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
RUN go build -o app

# Expose port 8080 to the outside world
EXPOSE 8080

# Start the application
CMD ["./app"]
