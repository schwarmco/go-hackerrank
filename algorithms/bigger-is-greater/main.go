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

	for i := 0; i < t; i++ {

		w := []byte(readLine(reader))

		pivot := len(w) - 1
		for pivot > 0 && w[pivot-1] >= w[pivot] {
			pivot--
		}

		if pivot <= 0 {
			fmt.Fprintln(writer, "no answer")
			continue
		}

		successor := len(w) - 1
		for w[successor] <= w[pivot-1] {
			successor--
		}

		w[pivot-1], w[successor] = w[successor], w[pivot-1]

		j := len(w) - 1
		for pivot < j {
			w[pivot], w[j] = w[j], w[pivot]
			pivot++
			j--
		}

		fmt.Fprintf(writer, "%s\n", w)
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
