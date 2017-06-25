BINARY=faker
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

build:
	${TESTS}
	go build -o ${BINARY}

install:
	${TESTS}
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)


clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
