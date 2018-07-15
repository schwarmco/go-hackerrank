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

	nt := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nt[0])
	t, _ := strconv.Atoi(nt[1])

	widths := make([]int, n)
	for idx, w := range strings.Split(readLine(reader), " ") {
		widths[idx], _ = strconv.Atoi(w)
	}

	for t > 0 {

		ij := strings.Split(readLine(reader), " ")
		i, _ := strconv.Atoi(ij[0])
		j, _ := strconv.Atoi(ij[1])

		minWidth := widths[i]
		for _, k := range widths[i : j+1] {
			if k < minWidth {
				minWidth = k
			}
		}

		fmt.Fprintln(writer, minWidth)

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
