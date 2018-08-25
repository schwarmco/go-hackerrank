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
	stocks := make([][]int, n)
	means := make([]float64, n)

	for i := 0; i < n; i++ {

		m, _ := strconv.Atoi(readLine(reader))
		stocks[i] = make([]int, m)

		var sum int
		for j, p := range strings.Split(readLine(reader), " ") {
			stocks[i][j], _ = strconv.Atoi(p)
			sum += stocks[i][j]
		}

		means[i] = float64(sum) / float64(len(stocks[i]))
	}

	means = unique(means)

	runningTimes := make([]float64, len(means))
	for idx, x := range means {
		for i := range stocks {
			for _, p := range stocks[i] {
				runningTimes[idx] += math.Abs(float64(p) - x)
			}
		}
	}

	sort.Float64s(runningTimes)

	fmt.Fprintf(writer, "%.12f", runningTimes[0])
}

func unique(input []float64) []float64 {
	u := make([]float64, 0, len(input))
	m := make(map[float64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
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
