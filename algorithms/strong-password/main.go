package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	numbers = "0123456789"
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = "!@#$%^&*()-+"
)

func solve(reader io.Reader, writer io.Writer) {

	readLine(reader)
	s := readLine(reader)

	var result int

	// need min 1 number
	if !strings.ContainsAny(s, numbers) {
		result++
	}

	// need min 1 special
	if !strings.ContainsAny(s, special) {
		result++
	}

	// need min 1 lower
	if !strings.ContainsAny(s, lower) {
		result++
	}

	// need min 1 upper
	if !strings.ContainsAny(s, upper) {
		result++
	}

	// need at least 6 chars
	if len(s)+result <= 6 {
		result = 6 - len(s)
	}

	fmt.Fprintf(writer, "%d\n", result)
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
