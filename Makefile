PROJECT_NAME=$(shell basename "$(PWD)")
SCRIPT_AUTHORS=Andrey Kapitonov <zikwall>
SCRIPT_VERSION=0.0.1.dev

.PHONY: lint
lint:
	golangci-lint run -c .golangci.yml

.PHONY: migrate-create
migrate-create:
	@read -p "Enter migration name:" name; \
	go run ./cmd/migrate/main.go --config-file ./config.yml create -n $$name

.PHONY: migrate-up
migrate-up:
	go run ./cmd/migrate/main.go --config-file ./config.yml up

.PHONY: migrate-down
migrate-down:
	go run ./cmd/migrate/main.go --config-file ./config.yml down