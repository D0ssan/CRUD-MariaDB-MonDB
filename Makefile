# Parameters
GOTOBIN=$(shell go env GOPATH)/bin

.PHONY: lint
lint:
	/bin/sh -c "if [ ! -f $(GOTOBIN)/golangci-lint ]; then \
 			curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOTOBIN) v1.42.1; \
    	fi"
	golangci-lint run

.PHONY: test
test:
	go test -v ./...

.PHONY: docker
docker:
	docker start my-mariadb