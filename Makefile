# Build the Go application
build:
	go build -o ./build/goShellCommander

# Run the built Go application
run:
	./build/goShellCommander

# Start the application using Docker Compose
up:
	docker-compose up -d

# Start the application using Docker Compose and rebuild the images
up-build:
	docker-compose up --build -d

# Run tests for the Go application
test:
	go test -v -cover ./...

# Declare the targets that are not associated with files
.PHONY: build run up up-build