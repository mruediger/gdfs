package gdfs
import (
	"crypto/sha1"
	"io"
	"fmt"
	"encoding/hex"
	"os"
	"github.com/rakyll/magicmime"
)

type gdfs_path string

type gdfs_hash []byte

type gdfs_mimetype string

type gdfs_file struct {
	path gdfs_path
	hash gdfs_hash
	mime gdfs_mimetype
}

func (hash gdfs_hash) String() string {
	return hex.EncodeToString(hash)
}

func (f gdfs_file) String() string {
	return fmt.Sprintf("{%s, %s, %s}", f.path, f.hash, f.mime )
}

func hash(path gdfs_path) (gdfs_hash, error) {
	f, err := os.Open(string(path))
	if err != nil {
		return nil,err
	}
	defer f.Close()
	h := sha1.New()
	io.Copy(h, f);
	return gdfs_hash(h.Sum(nil)),nil
}

func mime(path gdfs_path) (gdfs_mimetype, error) {
	mm, err := magicmime.New(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_ERROR)
	if err != nil {
		return gdfs_mimetype(""),err
	}
	mime,err := mm.TypeByFile(string(path))
	return gdfs_mimetype(mime),err
}

func New(path string) (file *gdfs_file , err error)  {
	file = new(gdfs_file)
	file.path = gdfs_path(path)
	file.hash,err = hash(file.path)
	file.mime,err = mime(file.path)
	return file, err
}
