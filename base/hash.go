package base

import (
	"crypto/sha256"
	"io"
	"log"
	"os"
)

// Returns the Sha256Sum of the file given by `path`.
func Sha256Sum(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return sha256Sum(f)
}

// Internal helper method for testing.
func sha256Sum(reader io.Reader) []byte {
	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}
