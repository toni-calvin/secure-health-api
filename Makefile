# Variables
PROJECT_NAME := topdoctors-challenge
APP_NAME := $(PROJECT_NAME)-app
TAG := latest
DOCKER_IMAGE := $(APP_NAME):$(TAG)
DOCKERFILE := Dockerfile
PORT := 8080
VOLUME_NAME := $(PROJECT_NAME)_postgres_data

# Default target: build and run
all: build run

# Build the Docker image without cache
build:
	@echo "Building Docker image..."
	docker build --no-cache -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

# Run the Docker container interactively
run:
	@echo "Running Docker container interactively..."
	docker run -it --rm -p $(PORT):$(PORT) $(DOCKER_IMAGE)

# Clean up only project-related resources
clean:
	@echo "Stopping and removing all containers, networks, and volumes for the project..."
	docker-compose down --volumes --remove-orphans
	@echo "Removing the application image..."
	-docker rmi $(DOCKER_IMAGE) || true
	@echo "Cleanup completed successfully."

# Full rebuild: clean, build, and run
rebuild: clean build run
