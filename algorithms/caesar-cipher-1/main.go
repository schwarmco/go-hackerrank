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

	n, _ := strconv.Atoi(readLine(reader))
	_ = n
	s := readLine(reader)
	k, _ := strconv.Atoi(readLine(reader))

	// if k > 26, we don't need to rotate more than that - just use rest
	k = k % 26

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	rotated, overflow := alphabet[:k], alphabet[k:]
	rotated = overflow + rotated

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			// use lowercase-offset to 0-indexed rotated -> 97
			fmt.Fprint(writer, string(rotated[c-97]))
		} else if c >= 'A' && c <= 'Z' {
			// use uppercase-offset to 0-indexed rotated -> 65
			// convert the result to uppercase -> -32
			fmt.Fprint(writer, string(rotated[c-65]-32))
		} else {
			fmt.Fprint(writer, string(c))
		}
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
