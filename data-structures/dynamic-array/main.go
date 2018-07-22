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

	nq := strings.Split(readLine(reader), " ")
	N, _ := strconv.Atoi(nq[0])
	Q, _ := strconv.Atoi(nq[1])

	seqList := make([][]int, N)
	var lastAnswer int

	for Q > 0 {

		txy := strings.Split(readLine(reader), " ")
		t, _ := strconv.Atoi(txy[0])
		x, _ := strconv.Atoi(txy[1])
		y, _ := strconv.Atoi(txy[2])

		idx := (x ^ lastAnswer) % N

		if t == 1 {
			seqList[idx] = append(seqList[idx], y)
		}

		if t == 2 {
			lastAnswer = seqList[idx][y%len(seqList[idx])]
			fmt.Fprintln(writer, lastAnswer)
		}

		Q--
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
