package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

// BigIntSlice implements the sort.Interface
type BigIntSlice []string

func (a BigIntSlice) Len() int      { return len(a) }
func (a BigIntSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BigIntSlice) Less(i, j int) bool {
	// we can maybe save bigint-conversion by just comparing lengths
	if len(a[i]) < len(a[j]) {
		return true
	}
	if len(a[i]) > len(a[j]) {
		return false
	}
	// if strings are equal length, we have to compare
	x, _ := new(big.Int).SetString(a[i], 10)
	y, _ := new(big.Int).SetString(a[j], 10)
	return x.Cmp(y) < 0
}

func solve(reader io.Reader, writer io.Writer) {

	n, _ := strconv.Atoi(readLine(reader))

	unsorted := make([]string, n)
	for i := 0; i < n; i++ {
		unsorted[i] = readLine(reader)
	}

	sort.Sort(BigIntSlice(unsorted))

	fmt.Fprintln(writer, strings.Join(unsorted, "\n"))
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
