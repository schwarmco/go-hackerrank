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

	a := make([]int, n)
	for idx, arr := range strings.Split(readLine(reader), " ") {
		a[idx], _ = strconv.Atoi(arr)
	}

	for len(a) > 0 {
		// print remaining sticks
		fmt.Fprintf(writer, "%d\n", len(a))
		// find smallest stick
		min, max := a[0], a[0]
		for _, i := range a {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		// discard if all sticks are the same length
		if min == max || min <= 0 {
			break
		}
		// remove sticks where there'd be nothing left
		i := len(a) - 1
		for i >= 0 {
			if a[i] <= min {
				if i == len(a)-1 {
					a = a[:i]
					i--
				} else {
					a = append(a[:i], a[i+1:]...)
				}
				continue
			}
			i--
		}
		// cut the remaining sticks by min
		for idx, i := range a {
			a[idx] = i - min
		}

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
