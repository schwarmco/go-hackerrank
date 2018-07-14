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

	// your code goes here
	p, _ := strconv.Atoi(readLine(reader))
	q, _ := strconv.Atoi(readLine(reader))

	numbers := []int{}

	for p <= q {

		n := math.Pow(float64(p), 2)

		nStr := strconv.FormatInt(int64(n), 10)
		d := int(math.Floor(float64(len(nStr)) / float64(2)))

		lStr, rStr := nStr[:int(d)], nStr[int(d):len(nStr)]

		l, _ := strconv.Atoi(lStr)
		r, _ := strconv.Atoi(rStr)

		if (l + r) == p {
			numbers = append(numbers, p)
		}

		p++
	}

	if len(numbers) > 0 {
		fmt.Fprintln(writer, strings.Trim(fmt.Sprint(numbers), "[]"))
	} else {
		fmt.Fprintln(writer, "INVALID RANGE")
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
