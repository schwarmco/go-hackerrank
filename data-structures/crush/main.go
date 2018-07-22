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

	nm := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	arr := make([]int, n+1)

	for m > 0 {

		abk := strings.Split(readLine(reader), " ")
		a, _ := strconv.Atoi(abk[0])
		b, _ := strconv.Atoi(abk[1])
		k, _ := strconv.Atoi(abk[2])

		arr[a] += k
		if b < n {
			arr[b+1] -= k
		}

		m--
	}

	var count, result int
	for i := 0; i <= n; i++ {
		count += arr[i]
		if count > result {
			result = count
		}
	}

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
