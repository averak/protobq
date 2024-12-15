.PHONY: install-tools
install-tools:
	go install ./cmd/protoc-gen-go-protobq
	go install github.com/bufbuild/buf/cmd/buf@v1.47.2
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: test
test:
	go test -v ./...

.PHONY: format
format:
	go fmt ./...
	golangci-lint run --issues-exit-code=0 --fix ./...
	buf format --write

BREAKING_CHANGE_BASE_BRANCH?=develop
.PHONY: lint
lint:
	golangci-lint run --issues-exit-code=1 ./...
	buf lint
	# buf breaking --against '.git#branch=$(BREAKING_CHANGE_BASE_BRANCH)'

.PHONY: codegen
codegen: install-tools
	find . -type f \( -name '*.pb.go' -or -name '*.protobq.go' \) -delete
	buf generate
