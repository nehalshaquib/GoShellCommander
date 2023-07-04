build:
	go build -o ./build/goShellCommander

run:
	./build/goShellCommander

up:
	docker-compose up -d

up-build:
	docker-compose up --build -d

.PHONY: build run up up-build