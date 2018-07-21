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

	q, _ := strconv.Atoi(readLine(reader))

	for q > 0 {

		s := readLine(reader)
		sDiff := distances(s)

		r := reverse(s)
		rDiff := distances(r)

		if equal(sDiff, rDiff) {
			fmt.Fprintln(writer, "Funny")
		} else {
			fmt.Fprintln(writer, "Not Funny")
		}

		q--
	}
}

func equal(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func distances(s string) (result []int) {
	for i := 0; i < len(s)-1; i++ {
		a := float64(s[i] - 97)
		b := float64(s[i+1] - 97)
		result = append(result, int(math.Abs(a-b)))
	}
	return
}

func reverse(s string) (result string) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
