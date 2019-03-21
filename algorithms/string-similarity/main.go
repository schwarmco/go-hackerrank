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

	for t > 0 {

		sum := 0
		s := readLine(reader)

		// using Z Algorithm instead
		var L, R int
		z := make([]int, len(s))
		for i := 1; i < len(s); i++ {
			if i > R {
				L = i
				R = i
				for R < len(s) && s[R-L] == s[R] {
					R++
				}
				z[i] = R - L
				R--
			} else {
				k := i - L
				if z[k] < R-i+1 {
					z[i] = z[k]
				} else {
					L = i
					for R < len(s) && s[R-L] == s[R] {
						R++
					}
					z[i] = R - L
					R--
				}
			}
		}

		for _, i := range z {
			sum += i
		}
		sum += len(s)
		fmt.Fprintf(writer, "%d\n", sum)

		// working but too slow:
		/*
			// iterating over each possible suffix
			for i := 0; i < len(s); i++ {

				// with prefixes upto suffix' length
				for j := 0; j < len(s)-i; j++ {
					// comparing prefix[i] with suffix[i]
					if s[j] != s[i:][j] {
						break
					}
					sum++
				}
			}
			fmt.Fprintf(writer, "%d\n", sum)
		*/

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
