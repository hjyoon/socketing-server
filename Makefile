GOCACHE ?= /tmp/socketing-go-cache
COVERPKG := ./internal/api,./internal/app,./internal/auth,./internal/routes

.PHONY: test coverage

test:
	GOCACHE=$(GOCACHE) go test -mod=mod ./...

coverage:
	GOCACHE=$(GOCACHE) go test -mod=mod ./internal/app ./internal/auth \
		-coverpkg=$(COVERPKG) -coverprofile=coverage.out
	GOCACHE=$(GOCACHE) go tool cover -func=coverage.out | tee /tmp/socketing-coverage.txt
	awk '/total:/ {gsub(/%/,"",$$3); if ($$3 < 97) exit 1}' /tmp/socketing-coverage.txt
