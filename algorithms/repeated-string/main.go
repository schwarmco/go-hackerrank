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
	s := readLine(reader)
	n, _ := strconv.Atoi(readLine(reader))

	// lets find how often the string is completely repeated and how much of
	// the remaining string we have to check
	x := n / len(s)
	y := math.Mod(float64(n), float64(len(s)))

	// count the occurances of "a" in the string
	count := 0
	for _, c := range s {
		if c == 'a' {
			count++
		}
	}

	// mulitply by x (how often the string is completely repeated)
	count = count * int(x)

	// finally, add the occurances in the remaining string
	for i := 0; i < int(y); i++ {
		if s[i] == 'a' {
			count++
		}
	}

	fmt.Fprintf(writer, "%d\n", count)
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
