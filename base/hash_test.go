package base

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSha256Sum(t *testing.T) {
	const str = "R29waGVycyBydWxlIQ=="
	const hash_for_str = "83fb9942fe4f65fae99f4db06df617533010a95131e7e26b67b34ad0a4ed83d3"

	buf := bytes.NewBufferString(str)
	hash := sha256Sum(buf)

	expected_hash, _ := hex.DecodeString(hash_for_str)
	if !reflect.DeepEqual(hash, expected_hash) {
		t.Fatalf("Expected hash %x, got %x", expected_hash, hash)
	}
}
