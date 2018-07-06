// bytecounter
package main

import "fmt"
import "io"
import "os"

type MyWriter struct {
	writer io.Writer
	point  *int64
}

func (w *MyWriter) Write(p []byte) (int, error) {
	*(w.point) = int64(len(p))
	w.writer.Write(p)

	return len(p), nil
}

func (myW *MyWriter) setWriter(writer io.Writer) {
	myW.writer = writer
}

func CountingWriter(writer io.Writer) (io.Writer, *int64) {
	myWriter := MyWriter{writer: writer}
	var tmp int64
	myWriter.point = &tmp
	return &myWriter, myWriter.point
}

func main() {
	writer, size := CountingWriter(os.Stdout)
	writer.Write([]byte("helloworld"))

	fmt.Printf("\n%d\n", *size)

}
