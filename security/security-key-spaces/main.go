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

	str := []int{}
	for _, s := range strings.Split(readLine(reader), "") {
		m, _ := strconv.Atoi(s)
		str = append(str, m)
	}

	e, _ := strconv.Atoi(readLine(reader))

	var result string
	for _, x := range str {
		result += strconv.Itoa(shift(x, e))
	}

	fmt.Fprintf(writer, "%s\n", result)
}

func shift(m, by int) (c int) {
	c = m + by
	for c > 9 {
		c -= 10
	}
	return
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
