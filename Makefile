BACKEND_BINARY=backendApp

up:
	@echo "Starting Docker images..."
	@docker-compose up -d
	@echo "Docker images started"

up_build: build_backend
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## build_backend: builds the broker binary as a linux executable
build_backend:
	@echo "Building broker binary..."
	cd ./cmd && env GOOS=linux CGO_ENABLED=0 go build -o ${BACKEND_BINARY} ./api
	@echo "Done!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"