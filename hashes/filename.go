package hashes

import (
	"crypto/sha1"
	"io/ioutil"
	"fmt"
)

func GetFilename(filename string) (out string, err error) {
	file,err := ioutil.ReadFile(filename)
	if err != nil {
		return "",err
	}

	h := sha1.New()
	h.Write(file)

	tmp := fmt.Sprintf("%x", h.Sum(nil))
	out = fmt.Sprintf("%s/%s", tmp[:2], tmp[2:])
	return out,err
}
