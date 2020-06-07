PACKAGES=github.com/xiejw/lunar/...

compile:
	go build ${PACKAGES}

fmt:
	go fmt ${PACKAGES}

test:
	go test ${PACKAGES}

bench:
	go test -bench=. ${PACKAGES}

