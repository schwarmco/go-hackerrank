package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func solve(reader io.Reader, writer io.Writer) {

	q, _ := strconv.Atoi(readLine(reader))

	for i := 0; i < q; i++ {

		am := strings.Split(readLine(reader), " ")
		a, _ := strconv.Atoi(am[0])
		m, _ := strconv.Atoi(am[1])

		var result int
		var wg sync.WaitGroup
		wg.Add(1)
		visit(&wg, m, a, math.MaxInt64, &result)
		wg.Wait()
		fmt.Fprintln(writer, result)
	}
}

func visit(wg *sync.WaitGroup, participant, shares, took int, result *int) {

	defer wg.Done()

	if participant < 0 {
		return
	}

	if shares == 0 {
		*result++
		return
	}

	x := int(math.Min(float64(shares), float64(took-1)))

	for i := 1; i <= x; i++ {
		wg.Add(1)
		go visit(wg, participant-1, shares-i, i, result)
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
