fmt:
	gofmt -w -l exec

test: fmt
	go test -v github.com/xiejw/lunar/...
