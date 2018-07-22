package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	n, _ := strconv.Atoi(readLine(reader))

	calories := make([]int, n)

	for idx, c := range strings.Split(readLine(reader), " ") {
		calories[idx], _ = strconv.Atoi(c)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	var result int
	for idx, c := range calories {
		result += int(math.Pow(2, float64(idx))) * c
	}

	fmt.Fprintln(writer, result)
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
