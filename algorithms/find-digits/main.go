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

		nStr := readLine(reader)
		n, _ := strconv.Atoi(nStr)

		var sum int
		for _, dStr := range nStr {
			d, _ := strconv.Atoi(string(dStr))
			if d != 0 && n%d == 0 {
				sum++
			}
		}

		fmt.Fprintf(writer, "%d\n", sum)

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
