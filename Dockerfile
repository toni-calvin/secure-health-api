# Use a base image with Go
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy all project files to the container
COPY . .

# Change directory to /app/cmd to build the Go application
WORKDIR /app/cmd

# Build the Go binary and place it in /app
RUN go build -o /app/main .

# Return to the root working directory
WORKDIR /app

# Ensure the binary is executable
RUN chmod +x /app/main

# Expose the port the application will listen on
EXPOSE 8080

# Run the application
CMD ["./main"]