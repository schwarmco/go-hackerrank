package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	readLine(reader)

	a := strings.Split(readLine(reader), " ")
	b := map[string]int{}

	for _, i := range a {
		b[i]++
	}

	for s, i := range b {
		if i == 1 {
			fmt.Fprintln(writer, s)
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
