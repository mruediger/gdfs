package gdfs

import (
	"testing"
	"fmt"
)


func TestReadFile(t *testing.T) {
	file,err := New("gdfs_file_test.go")
	if (err != nil) {
		panic(err)
	}
	fmt.Printf("%v", file)
}

func TestMimetype(t *testing.T) {
	m,_ := mime("gdfs_file.go")
	equals(string(m) , "text/plain", t)
}

func TestHash(t *testing.T) {
	h,_ := hash("gdfs_file.go")
	equals(h.String(), "e256612f5548b0837be2ffa5ce55958b746f56e3", t)
}

func equals(got, wanted interface{}, t *testing.T) {
	if got != wanted {
		t.Errorf("wanted %q, got %q", wanted, got)
	}
}