build:
	@go build -o bin/GoLang-Postgres-Auth0

run: build
	@./bin/GoLang-Postgres-Auth0

test:
	@go test -v ./...

