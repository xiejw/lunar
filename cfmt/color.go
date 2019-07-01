package cfmt

import (
	"fmt"
)

const (
	red     = "\x1b[31m"
	green   = "\x1b[32m"
	yellow  = "\x1b[33m"
	blue    = "\x1b[34m"
	magenta = "\x1b[35m"
	cyan    = "\x1b[36m"
)

const (
	reset = "\x1b[0m"
)

const colorFormat = "%v%v%v"

// An alias to allow tests swapping the Printf to avoid real I/O.
var printf = fmt.Printf

func colorPrint(color string, format string, args ...interface{}) {
	stringToPrint := fmt.Sprintf(format, args...)
	printf(colorFormat, color, stringToPrint, reset)
}

func Red(format string, args ...interface{})     { colorPrint(red, format, args...) }
func Green(format string, args ...interface{})   { colorPrint(green, format, args...) }
func Yellow(format string, args ...interface{})  { colorPrint(yellow, format, args...) }
func Blue(format string, args ...interface{})    { colorPrint(blue, format, args...) }
func Magenta(format string, args ...interface{}) { colorPrint(magenta, format, args...) }
func Cyan(format string, args ...interface{})    { colorPrint(cyan, format, args...) }
