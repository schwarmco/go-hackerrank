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

	l, _ := strconv.Atoi(readLine(reader))
	r, _ := strconv.Atoi(readLine(reader))

	arr := [][]int{}
	for i := l; i <= r; i++ {
		for j := l; j <= r; j++ {
			arr = append(arr, []int{i, j})
		}
	}

	var max int
	for _, a := range arr {
		r := a[0] ^ a[1]
		if r > max {
			max = r
		}
	}

	fmt.Fprintln(writer, max)
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
