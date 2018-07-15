package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func solve(reader io.Reader, writer io.Writer) {

	ret, _ := time.Parse("2 1 2006", readLine(reader))
	due, _ := time.Parse("2 1 2006", readLine(reader))

	var fine int
	if ret.Before(due) {
		fine = 0
	} else if ret.Year() > due.Year() {
		fine = 10000
	} else if ret.Month() > due.Month() {
		fine = int(ret.Month()-due.Month()) * 500
	} else if ret.Day() > due.Day() {
		fine = int(ret.Day()-due.Day()) * 15
	}

	fmt.Fprintln(writer, fine)
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
