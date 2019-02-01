BINARY=faker

build: test
	@go build ./...

test: 
	@go test -v ./...

unittest:
	@go test -v -short ./...

# Linter
lint-prepare: 
	@echo "Installing golangci-lint"
	@go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

lint: 
	golangci-lint run \
		--exclude="cyclomatic complexity" \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: lint lint-prepare clean build unittest 