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
	n, _ := strconv.Atoi(readLine(reader))

	// instead of saving the value to the index, we use the index as value, so
	// we can just use x (referring to the value) as the index later one
	p := make([]int, n+1)
	for idx, value := range strings.Split(readLine(reader), " ") {
		v, _ := strconv.Atoi(value)
		p[v] = idx + 1
	}
	fmt.Println(p)

	for x := 1; x <= n; x++ {
		fmt.Fprintln(writer, p[p[x]])
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
