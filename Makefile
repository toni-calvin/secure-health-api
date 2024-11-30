# Variables
PROJECT_NAME := topdoctors-challenge
APP_NAME := $(PROJECT_NAME)-app
TAG := latest
DOCKER_IMAGE := $(APP_NAME):$(TAG)
PORT := 8080

# Default target: build and run
all: build run

# Build the Docker image using docker-compose
build:
	@echo "Building Docker images with docker-compose..."
	docker-compose build

# Run the application using docker-compose
run:
	@echo "Starting the application using docker-compose..."
	docker-compose up

# Stop the application without removing volumes
stop:
	@echo "Stopping all containers..."
	docker-compose stop

# Clean up all project resources, including volumes
clean:
	@echo "Stopping and removing all containers, networks, and volumes for the project..."
	docker-compose down --volumes --remove-orphans
	@echo "Removing the application image..."
	-docker rmi $(DOCKER_IMAGE) || true
	@echo "Cleanup completed successfully."

# Full rebuild: clean and run using docker-compose
rebuild: clean build run
