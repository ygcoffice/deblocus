package geo

import (
	"bytes"
	"compress/zlib"
	"io"
)

func buildGeoDB() []byte{
	var db = []byte(
	return decompress(db)
}

func decompress(b []byte) []byte {
	r := bytes.NewReader(b)
	zr , e := zlib.NewReader(r)
	if e != nil {
		panic(e)
	}
	w := new(bytes.Buffer)
	io.Copy(w, zr)
	return w.Bytes()
}