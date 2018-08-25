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

	var c string
	for _, s := range strings.Split(readLine(reader), "") {
		m, _ := strconv.Atoi(s)
		c += strconv.Itoa(shift(m))
	}

	fmt.Fprintf(writer, "%s\n", c)
}

func shift(m int) int {
	if m >= 9 {
		return 0
	}
	return m + 1
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
