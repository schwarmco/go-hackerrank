package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	p, _ := strconv.Atoi(readLine(reader))

	for p > 0 {

		s1 := readLine(reader)
		s2 := readLine(reader)

		r := regexp.MustCompile("[" + s1 + "]+")
		if r.MatchString(s2) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}

		p--
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
