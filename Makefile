BUILD_FILENAME = ova-obligation-api

test:
	go test ./...

build:
	go get ./...
	go build -o $(BUILD_FILENAME) ./cmd/ova-obligation-api/main.go
