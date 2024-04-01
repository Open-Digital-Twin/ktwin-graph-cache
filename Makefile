build:
	go build -o ./

run-local:
	export LOCAL=true && \
	go run .

run:
	go run .

swag:
	swag init

wire:
	wire ./internal/app && \
	wire ./internal/app/context/...
	
docker-build-push:
	docker build -t ghcr.io/open-digital-twin/ktwin-graph-store:0.1 . && docker push ghcr.io/open-digital-twin/ktwin-graph-store:0.1