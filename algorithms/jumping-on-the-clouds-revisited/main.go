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

	nk := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	c := make([]int, n)
	for idx, num := range strings.Split(readLine(reader), " ") {
		c[idx], _ = strconv.Atoi(num)
	}

	e := 100
	i := 0
	for {
		i = (i + k) % n
		e--
		if c[i] == 1 {
			e -= 2
		}
		if i == 0 {
			break
		}
	}

	fmt.Fprintln(writer, e)
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
