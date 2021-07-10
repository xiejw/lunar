package base

import (
	"flag"

	"github.com/xiejw/lunar/glog"
)

// Initializes the lunar.
func Init(parseFlag bool) {
	if parseFlag {
		flag.Parse()
	}
}

// Finalizes the lunar.
func FinishUp() {
	glog.Flush()
}
