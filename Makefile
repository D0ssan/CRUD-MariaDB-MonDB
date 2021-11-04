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

.PHONY: mariadb
mariadb:
#!/bin/bash
	#!/bin/bash
	sudo mysql -e "CREATE DATABASE IF NOT EXISTS test_users;" -u root
	export PATH=$PATH:$HOME/GOPATH/bin

	set -ex

	migrate -path databases/mariadb/migration -database "mysql://root:ppasword@tcp(localhost:3306)/test_users" up

