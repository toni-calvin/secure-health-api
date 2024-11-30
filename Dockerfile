# Use Go base image
FROM golang:1.23.3-alpine

# Set the working directory
WORKDIR /app

# Copy the project files
COPY . .

# Navigate to the cmd directory to build the Go binary
WORKDIR /app/cmd

# Build the Go binary and place it in /app
RUN go build -o /app/main .

# Return to the root working directory
WORKDIR /app

# Ensure the binary is executable
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./main"]