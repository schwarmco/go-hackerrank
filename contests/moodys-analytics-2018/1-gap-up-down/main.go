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

	// input

	n, _ := strconv.Atoi(readLine(reader))
	L := make([]int, n)
	for i, low := range strings.Split(readLine(reader), " ") {
		L[i], _ = strconv.Atoi(low)
	}
	H := make([]int, n)
	for i, high := range strings.Split(readLine(reader), " ") {
		H[i], _ = strconv.Atoi(high)
	}
	C := make([]int, n)
	for i, close := range strings.Split(readLine(reader), " ") {
		C[i], _ = strconv.Atoi(close)
	}

	// process

	var gapUp, gapDown int

	for i := 1; i < n; i++ {
		if C[i-1] < L[i] {
			gapUp++
		}
		if C[i-1] > H[i] {
			gapDown++
		}
	}

	// output

	fmt.Fprintf(writer, "%d %d\n", gapUp, gapDown)
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
