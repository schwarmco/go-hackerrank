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

	t, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < t; i++ {

		nms := strings.Split(readLine(reader), " ")
		n, _ := strconv.Atoi(nms[0])
		m, _ := strconv.Atoi(nms[1])
		s, _ := strconv.Atoi(nms[2])

		fmt.Fprintf(writer, "%d\n", ((m-1)+(s-1))%n+1)
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
