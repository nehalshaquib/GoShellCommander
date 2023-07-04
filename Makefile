build:
	go build -o ./build/goShellCommander

run:
	./build/goShellCommander

up:
	docker-compose up -d

up-build:
	docker-compose up --build -d

test:
	go test -v -cover ./...

.PHONY: build run up up-build