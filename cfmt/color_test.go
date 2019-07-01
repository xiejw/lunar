package cfmt

import (
	"testing"
)

func dummyPrintf(format string, args ...interface{}) (int, error) {
	return len(format), nil
}

func TestRed(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Red("Hello")
}

func TestGreen(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Green("Hello")
}

func TestYellow(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Yellow("Hello")
}

func TestBlue(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Blue("Hello")
}

func TestMagenta(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Magenta("Hello")
}

func TestCyan(t *testing.T) {
	oldPrintf := printf
	printf = dummyPrintf
	defer func() {
		printf = oldPrintf
	}()

	Cyan("Hello")
}
