package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func solve(reader io.Reader, writer io.Writer) {

	// your code goes here
	nm := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])

	// using bigInts, because length of a bit-string may be up to 500
	attendees := make([]big.Int, n)
	for i := 0; i < n; i++ {
		attendees[i].SetString(readLine(reader), 2)
	}

	// index = num of topics, value = num of teams
	topics := make([]int, m+1)
	for aIdx, aTopics := range attendees {
		for bIdx, bTopics := range attendees {
			if aIdx == bIdx {
				continue
			}
			// calculate to OR
			or := new(big.Int).Or(&aTopics, &bTopics)
			// count the 1's
			numTopics := 0
			for _, v := range or.Bits() {
				numTopics += bits.OnesCount64(uint64(v))
			}
			topics[numTopics]++
		}
	}

	// find the highest non-zero number of topcis
	for i := m; i > 0; i-- {
		if topics[i] != 0 {
			fmt.Fprintln(writer, i)
			// divide it by 2, to get rid of doubled teams (0,1 - 1,0)
			fmt.Fprintln(writer, topics[i]/2)
			break
		}
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
