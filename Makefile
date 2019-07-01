fmt:
	gofmt -w -l .

test: fmt
	go test -v github.com/xiejw/lunar/...
