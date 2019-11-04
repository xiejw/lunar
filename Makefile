BUILD=build
FMT=gofmt -w -l

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: fmt

fmt:
	${FMT} .

test: fmt
	go test ${TEST_VERBOSE} github.com/xiejw/lunar/...
