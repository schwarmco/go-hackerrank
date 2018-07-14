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

	M := [][]int{}
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for idx, r := range strings.Split(readLine(reader), "") {
			row[idx], _ = strconv.Atoi(r)
		}
		M = append(M, row)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < 1 || j < 1 || i >= n-1 || j >= n-1 {
				fmt.Fprint(writer, M[i][j])
				continue
			}
			cur := M[i][j]
			top, bottom := M[i-1][j], M[i+1][j]
			right, left := M[i][j+1], M[i][j-1]
			if cur > top && cur > bottom && cur > right && cur > left {
				fmt.Fprint(writer, "X")
			} else {
				fmt.Fprint(writer, cur)
			}
		}
		fmt.Fprintln(writer, "")
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
