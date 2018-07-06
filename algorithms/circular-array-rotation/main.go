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

	// your code goes here
	nkq := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nkq[0])
	k, _ := strconv.Atoi(nkq[1])
	q, _ := strconv.Atoi(nkq[2])

	// read and make integer-slice
	a := make([]int, n)
	for idx, arr := range strings.Split(readLine(reader), " ") {
		a[idx], _ = strconv.Atoi(arr)
	}

	// rotate a k times
	// --- actually the problem requires you to solves this with math
	// --- although i'm just failing Test-Case #4 due to timeout
	// x, a := a[len(a)-k:], a[:len(a)-k]
	// a = append(x, a...)

	shift := n - (k % n)

	// output requested values
	for q > 0 {
		m, _ := strconv.Atoi(readLine(reader))
		// fmt.Fprintf(writer, "%d\n", a[m])
		idx := (shift + m) % n
		fmt.Fprintf(writer, "%d\n", a[idx])
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
