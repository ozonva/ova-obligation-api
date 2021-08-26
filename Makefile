BUILD_FILENAME = ova-obligation-api

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