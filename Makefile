BUILD=build
FMT=gofmt -w -l

ifdef VERBOSE
	TEST_VERBOSE=-v
endif

default: fmt compile test

compile:
	go build -o ${BUILD}/dependency_viewer cmd/dependency_viewer/main.go


fmt:
	${FMT} .

test: fmt
	go test ${TEST_VERBOSE} github.com/xiejw/lunar/...
