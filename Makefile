# 		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.1

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	MYMARIADB_PASSWORD=secret go test -v ./...
