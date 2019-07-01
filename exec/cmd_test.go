package exec

import (
	"os"
	"testing"
)

func TestRunCMD(t *testing.T) {
	results, err := RunCmd("whoami")
	if err != nil {
		t.Fatalf("Not expect error.")
	}
	expected := os.Getenv("USER")
	if results[0] != expected {
		t.Fatalf("got %s, expected %s.", results[0], expected)
	}
}
