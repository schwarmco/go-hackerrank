package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	// your code goes here
	s := readLine(reader)

	msg := [3]rune{'S', 'O', 'S'}
	count := 0

	for idx, c := range s {
		if c != msg[idx%3] {
			count++
		}
	}

	fmt.Fprintf(writer, "%d\n", count)
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		panic(err)
	}
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	solve(reader, writer)

	writer.Flush()
}

func readLine(reader io.Reader) string {
	r := bufio.NewReader(reader)
	str, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
