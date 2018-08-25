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

	q, _ := strconv.ParseInt(readLine(reader), 10, 64)

	for i := int64(0); i < q; i++ {

		n, _ := strconv.ParseInt(readLine(reader), 10, 64)

		var carryover int64

		for i := int64(0); i < n; i++ {

			ab := strings.Split(readLine(reader), " ")
			actual, _ := strconv.ParseInt(ab[0], 10, 64)
			estimated, _ := strconv.ParseInt(ab[1], 10, 64)
			estimated += carryover

			carryover = 0
			if estimated > actual {
				carryover = estimated - actual
			}
		}

		if carryover > 0 {
			fmt.Fprintln(writer, "1")
		} else {
			fmt.Fprintln(writer, "0")
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
