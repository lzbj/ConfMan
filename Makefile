BINARY=ConfMan
TESTS=go test $$(go list ./... | grep -v /vendor/) -cover

build:
	go build -o ${BINARY}

install:
	${TESTS}
	go build -o ${BINARY}

unittest:
	go test -short $$(go list ./... | grep -v /vendor/)


clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

fmt:
	find . -name "*.go" -not -path "./vendor/*" -not -path ".git/*" | xargs gofmt -s -d -w

.PHONY: clean install unittest format
