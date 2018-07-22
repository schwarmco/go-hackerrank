package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	q, _ := strconv.Atoi(readLine(reader))

	for q > 0 {
		n, _ := strconv.ParseUint(readLine(reader), 10, 32)
		fmt.Fprintln(writer, uint32(^n))
		q--
	}
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
