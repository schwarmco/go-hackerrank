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

	n, _ := strconv.Atoi(readLine(reader))

	// the problems function
	f := func(x int) int {
		return x
	}

	X := make([]int, n)
	Y := make([]int, n)
	for idx, x := range strings.Split(readLine(reader), " ") {
		X[idx], _ = strconv.Atoi(x)
		Y[idx] = f(X[idx])
	}

	if isBijective(X, Y) {
		fmt.Fprintln(writer, "YES")
	} else {
		fmt.Fprintln(writer, "NO")
	}
}

func isBijective(X, Y []int) bool {
	if len(X) != len(Y) {
		return false
	}
	return isUnique(X) && isUnique(Y)
}

func isUnique(X []int) bool {
	m := make(map[int]bool)
	for _, x := range X {
		m[x] = true
	}
	return len(m) == len(X)
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
