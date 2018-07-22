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

	t, _ := strconv.Atoi(readLine(reader))

	for t > 0 {

		y, _ := strconv.Atoi(readLine(reader))
		z := y

		for z%3 != 0 {
			z -= 5
		}

		if z < 0 {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, strings.Repeat("5", z)+strings.Repeat("3", y-z))
		}

		t--
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
