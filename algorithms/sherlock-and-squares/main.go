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
	t, _ := strconv.Atoi(readLine(reader))

	for t > 0 {

		ab := strings.Split(readLine(reader), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])

		// we need to find the first square number >= a
		square := 0
		for square <= 0 {
			tmp := int(math.Sqrt(float64(a)))
			if (tmp * tmp) == a {
				square = a
				break
			}
			a++
		}

		count := 0
		for square <= b {
			distance := (2 * int(math.Sqrt(float64(square)))) + 1
			square += distance
			count++
		}

		fmt.Fprintf(writer, "%d\n", count)

		t--
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
