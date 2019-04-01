BIN_NAME=flaw
BIN_OUTPUT=dist/${BIN_NAME}

fmt:
	go fmt ./...

deps:
	go mod vendor
	go mod verify

test:
	go test ./...

build: fmt deps
	go build -o ${BIN_OUTPUT}
