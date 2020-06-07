package lunar

import (
	"flag"

	"github.com/golang/glog"
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
