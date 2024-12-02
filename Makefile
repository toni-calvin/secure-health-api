.PHONY: all build run stop clean rebuild database

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

# Run the app 
run:
	@echo "Starting the application using docker-compose..."
	docker-compose up app db

# Run tests
test:
	@echo "Running tests..."
	docker-compose up -d test_db
	@docker-compose exec test_db sh -c 'until pg_isready -U postgres; do sleep 1; done;'
	docker-compose run --rm test
	docker-compose down



stop:
	@echo "Stopping all containers..."
	docker-compose stop

clean:
	@echo "Stopping and removing all containers, networks, and volumes for the project..."
	docker-compose down --volumes --remove-orphans
	@echo "Removing the application image..."
	-docker rmi $(DOCKER_IMAGE) || true
	@echo "Pruning unused Docker volumes and images..."
	docker volume prune -f
	docker image prune -f
	@echo "Cleanup completed successfully."


# Full rebuild: clean and run using docker-compose
rebuild: clean build run

database:
	@echo "Starting the database using docker-compose..."
	docker-compose exec db psql -U postgres -d topdoctors