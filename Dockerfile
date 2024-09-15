# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install make
RUN apk add --no-cache make

# Build the Go app
RUN go build -o bin/proj main.go

# Expose port 3000 to the outside world
EXPOSE 3000

# Export all environment variables to a .env.local file
RUN printenv > .env.local

# Command to run the executable
CMD ["./bin/proj"]