BINARY=faker
TESTS=go test -race -coverprofile=coverage.txt -covermode=atomic


build:
	go build -o ${BINARY}

test:
	${TESTS}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)


clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install unittest
