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
	arr := make([]int, n)
	for idx, num := range strings.Split(readLine(reader), " ") {
		arr[idx], _ = strconv.Atoi(num)
	}

	value := arr[len(arr)-1]

	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
		if arr[i] < value {
			arr[i] = value
			fmt.Fprintln(writer, strings.Trim(fmt.Sprint(arr), "[]"))
			break
		}
		fmt.Fprintln(writer, strings.Trim(fmt.Sprint(arr), "[]"))
	}

	if arr[0] > value {
		arr[0] = value
		fmt.Fprintln(writer, strings.Trim(fmt.Sprint(arr), "[]"))
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
