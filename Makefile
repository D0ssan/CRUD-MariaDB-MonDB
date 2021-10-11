GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOLANGCI=go-lint

# MYMARIADB_PASSWORD=secret go test -cover github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb github.com/d0ssan/CRUD-MariaDB-MongoDB/configs
# MYMARIADB_PASSWORD=secret $(GOTEST) ./...*_test.go -v
.PHONY: test
test:
	MYMARIADB_PASSWORD=secret go test github.com/d0ssan/CRUD-MariaDB-MongoDB/databases/mariadb github.com/d0ssan/CRUD-MariaDB-MongoDB/configs -v


.PHONY: lint
lint:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0
	$(GOLANGCI) run