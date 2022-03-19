.PHONY: test clean tidy deps \
	lint lint-go

test:
	go test ./...

tidy:
	go mod tidy

deps:
	go install github.com/mgechev/revive@latest
	go install golang.org/x/tools/cmd/goimports@latest

lint: lint-go

lint-go:
	revive -formatter stylish -config revive.toml ./...
