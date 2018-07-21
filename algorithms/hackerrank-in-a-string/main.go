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

	q, _ := strconv.Atoi(readLine(reader))

	r := regexp.MustCompile("([h]+.*[a]+.*[c]+.*[k]+.*[e]+.*[r]+.*[r]+.*[a]+.*[n]+.*[k]+.*)")

	for q > 0 {

		s := readLine(reader)
		if r.MatchString(s) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}

		q--
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
