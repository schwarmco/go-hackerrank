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

func solve(reader io.Reader, writer io.Writer) {

	n, _ := strconv.Atoi(readLine(reader))
	arr := make([]int, n)

	for idx, val := range strings.Split(readLine(reader), " ") {
		arr[idx], _ = strconv.Atoi(val)
	}

	sort.Ints(arr)

	smallestDistance := arr[1] - arr[0]
	d := map[int][]string{}
	for i := 0; i < len(arr)-1; i++ {
		distance := arr[i+1] - arr[i]
		if distance > smallestDistance {
			continue
		}
		d[distance] = append(d[distance], fmt.Sprint(arr[i], " ", arr[i+1]))
		smallestDistance = distance
	}

	fmt.Fprintln(writer, strings.Join(d[smallestDistance], " "))
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
