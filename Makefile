BUILD_FILENAME = ova-obligation-api
DB_PASSWORD = $(shell cat config/.env | grep POSTGRES_PASSWORD= | cut -d '=' -f2)
DB_USER = $(shell cat config/.env | grep POSTGRES_USER= | cut -d '=' -f2)
DB_NAME = $(shell cat config/.env | grep POSTGRES_DB= | cut -d '=' -f2)
GOBIN = $(shell go env | grep GOPATH= | cut -d '=' -f2)/bin

.PHONY: migration-status
migration-status:
	${GOBIN}/goose -dir=migration postgres "user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable" status

.PHONY: migration-create
migration-create:
	${GOBIN}/goose -dir=migration postgres "user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable" \
	create ${NAME} sql

.PHONY: migration-up
migration-up:
	${GOBIN}/goose -dir=migration postgres "user=${DB_USER} dbname=${DB_NAME} password=${DB_PASSWORD} sslmode=disable" \
	up

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go get ./...
	go build -o $(BUILD_FILENAME) ./cmd/ova-obligation-api/main.go

.PHONY: generate_proto
generate_proto:
	mkdir -p pkg/ova-obligation-api
	protoc \
			--go_out=pkg/ova-obligation-api --go_opt=paths=import \
			--go-grpc_out=pkg/ova-obligation-api --go-grpc_opt=paths=import \
			api/ova-obligation-api.proto
	mv pkg/ova-obligation-api/github.com/ozonva/ova-obligation-api/pkg/ova-obligation-api/* pkg/ova-obligation-api/
	rm -rf pkg/ova-obligation-api/github.com