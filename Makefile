build:
	@go build -o bin/api
run: build
	@./bin/api
test:
	@go test -v ./...
mongodb:
	@docker-compose up --build