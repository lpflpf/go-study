// 7.5
package main

import (
	"fmt"
	"io"
	"os"
)

type MyReader struct {
	limit  int64
	reader io.Reader
}

func (r *MyReader) Read(p []byte) (int, error) {
	n := int(r.limit)
	if len(p) < n {
		n = len(p)
	}
	r.reader.Read(p[0:n])
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	myReader := MyReader{}
	myReader.limit = n
	myReader.reader = r
	return &myReader
}

func main() {
	r := LimitReader(os.Stdin, 10)
	n, _ := r.Read([]byte("asdfa"))

	fmt.Println(n)
}
