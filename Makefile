build:
	@go build -o bin/expensebuddy src/main.go

run: build
	@./bin/expensebuddy

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir src/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run src/migrate/main.go up

migrate-down:
	@go run src/migrate/main.go down