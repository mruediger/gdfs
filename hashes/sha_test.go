package hashes

import (
	"crypto/sha1"
	"testing"
	"fmt"
	"io"
)

func TestSHA1(t *testing.T) {
	h := sha1.New()
	io.WriteString(h, "kartoffelbrei")

	out := fmt.Sprintf("%x", h.Sum(nil))

	if out != "b5dbe45cd8122a3ccd71a714b004faa3ba0a14ae" {
		t.Errorf("%s", out)
	}

}


func TestHashToFilename(t *testing.T) {
	hash := "b5dbe45cd8122a3ccd71a714b004faa3ba0a14ae"
	out := fmt.Sprintf("%s/%s", hash[:2], hash[2:])

	if out != "b5/dbe45cd8122a3ccd71a714b004faa3ba0a14ae" {
		t.Errorf("%s", out)
	}
}

func TestFilename(t *testing.T) {
	hash,_ := GetFilename("filename.go")

	if hash != "99/b0384824240fc922811c711b97ad4a961d4d56" {
		t.Errorf("%s\n", hash)
	}
}
