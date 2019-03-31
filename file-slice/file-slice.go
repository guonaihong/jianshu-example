package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func sliceFile(rs io.ReadSeeker, w io.Writer, start, end int) error {
	if start < 0 {
		return fmt.Errorf("invalid start argument: %d", start)
	}

	if end < 0 {
		return fmt.Errorf("invalid end argument: %d", end)
	}

	if end < start {
		return fmt.Errorf("invalid start:end argument: [%d, %d]", start, end)
	}

	_, err := rs.Seek(int64(start), 0)
	if err != nil {
		return err
	}

	canRead := end - start

	// buf大小可以调整，为了测试方便调小一点
	readBuf := make([]byte, 8)
	for canRead > 0 {
		needRead := len(readBuf)
		if needRead > canRead {
			needRead = canRead
		}

		n, err := rs.Read(readBuf[:needRead])
		if err != nil {
			break
		}
		w.Write(readBuf[:n])
		canRead -= n
	}
	return nil
}

func main() {

	start := flag.Int("s", 0, "start")
	end := flag.Int("e", 0, "end")
	flag.Parse()

	args := flag.Args()

	for _, v := range args {
		fd, err := os.Open(v)
		if err != nil {
		}

		sliceFile(fd, os.Stdout, *start, *end)
		fd.Close()
	}
}
