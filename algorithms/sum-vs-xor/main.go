package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	n, _ := strconv.Atoi(readLine(reader))

	var result int

	/*
		for i := 0; i <= n; i++ {
			if i+n == i^n {
				result++
			}
		}
	*/

	// due to timeouts of testcases 7-11,
	// this problem requires a more mathematical solution
	for n > 0 {
		if n%2 == 0 {
			result++
		}
		n /= 2
	}
	result = int(math.Pow(2, float64(result)))

	fmt.Fprintln(writer, result)
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
