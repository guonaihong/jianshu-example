package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func errorReadLine1(r io.Reader) {
	br := bufio.NewReader(r)
	for {
		l, e := br.ReadBytes('\n')
		if e == io.EOF {
			break
		}
		os.Stdout.Write(l)
	}
}

func errorReadLine2(r io.Reader) {
	br := bufio.NewReader(r)
	for {
		l, e := br.ReadBytes('\n')
		if e != nil {
			break
		}
		os.Stdout.Write(l)
	}
}

func readLine(r io.Reader) {
	br := bufio.NewReader(r)
	for {
		l, e := br.ReadBytes('\n')
		if e != nil && len(l) == 0 {
			break
		}
		os.Stdout.Write(l)
	}
}

func main() {
	sr := strings.NewReader("123\n456")

	fmt.Printf("start run errorReadLine1\n")
	errorReadLine1(sr)
	sr.Seek(0, 0)

	fmt.Printf("start run errorReadLine2\n")
	errorReadLine2(sr)
	sr.Seek(0, 0)

	fmt.Printf("start run readLine\n")
	readLine(sr)
}
