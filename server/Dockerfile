# Use the official Go image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app
COPY . /app

# Download and cache dependencies
RUN go mod download

# Build the Go application
RUN go build -o server ./main

# Set the entry point for the container
CMD ["./server"]
