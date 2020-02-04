package main

import (
	"io"
	"os"
	"strconv"
)

// Counter ...
type Counter struct {
	Writer io.Writer
}

// Count ...
func (c *Counter) Count(n uint64) uint64 {

	if n == 0 {
		c.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	c.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return c.Count(n - 1)
}

func main() {
	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()
	var c Counter = Counter{
		Writer: pipeWriter,
	}
	file, _ := os.Create("./test")
	defer file.Close()
	tee := io.TeeReader(pipeReader, file)
	go func() {
		io.Copy(os.Stdout, tee)
	}()
	c.Count(5)
}
