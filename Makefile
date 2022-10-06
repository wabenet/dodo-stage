.PHONY: all
all: clean test build lint

.PHONY: clean
clean:
	rm -rf ./dist

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: update
update:
	go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all | xargs --no-run-if-empty go get
	go mod tidy

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: api
	go test -cover -race ./...

.PHONY: build
build: api
	goreleaser build --snapshot --rm-dist

.PHONY: api
build: $(shell ls api/**/*.proto | sed 's|\.proto|.pb.go|g' | xargs)

%.pb.go: %.proto
	protoc -I . --go_out=plugins=grpc:. --go_opt=module=github.com/wabenet/dodo-stage $<
