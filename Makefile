GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1

build:
	go build ./...

run:
	go run ./... -c ./config/config.yaml

test:
	go test ./... -v -cover

lint:
	$(GOLANGCI_LINT) run --fix