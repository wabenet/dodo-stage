all: clean test build

.PHONY: clean
clean:
	rm -f dodo-stage_*

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run --enable-all

.PHONY: test
test: pkg/types/types.pb.go
	go generate ./...
	go test -cover ./...

.PHONY: build
build: pkg/types/types.pb.go
	go generate ./...
	gox -arch="amd64" -os="darwin linux" ./...

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. --go_opt=module=github.com/dodo-cli/dodo-stage $<
