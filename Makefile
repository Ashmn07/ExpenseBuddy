build:
	@go build -o bin/expensebuddy src/main.go

run: build
	@./bin/expensebuddy

test:
	@go test -v ./...