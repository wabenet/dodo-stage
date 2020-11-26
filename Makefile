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
test: api/v1alpha1/stage_plugin.pb.go api/v1alpha1/stage.pb.go
	go generate ./...
	go test -cover ./...

.PHONY: build
build: api/v1alpha1/stage_plugin.pb.go api/v1alpha1/stage.pb.go
	go generate ./...
	gox -arch="amd64" -os="darwin linux" ./...

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. --go_opt=module=github.com/dodo-cli/dodo-stage $<
