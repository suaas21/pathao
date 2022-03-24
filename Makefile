build:
	./build.sh

run: build
	go run ./cmd/pathao/main.go ./cmd/pathao/graphql.go serve --config config.yaml

serve:
	docker-compose down
	docker-compose up -d
