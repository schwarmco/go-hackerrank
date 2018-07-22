package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type customer struct {
	ID          int
	OrderNumber int
	PrepTime    int
	ServeTime   int
}

func solve(reader io.Reader, writer io.Writer) {

	n, _ := strconv.Atoi(readLine(reader))

	c := make([]customer, n)

	for i := 0; i < n; i++ {
		op := strings.Split(readLine(reader), " ")
		o, _ := strconv.Atoi(op[0])
		p, _ := strconv.Atoi(op[1])
		c[i] = customer{i + 1, o, p, o + p}
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].ServeTime < c[j].ServeTime
	})

	var result []string
	for _, cust := range c {
		result = append(result, fmt.Sprint(cust.ID))
	}

	fmt.Fprintln(writer, strings.Join(result, " "))
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
