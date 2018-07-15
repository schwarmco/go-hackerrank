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

	pdms := strings.Split(readLine(reader), " ")
	p, _ := strconv.Atoi(pdms[0])
	d, _ := strconv.Atoi(pdms[1])
	m, _ := strconv.Atoi(pdms[2])
	s, _ := strconv.Atoi(pdms[3])

	numGames := 0
	for p > m && s-p > 0 {
		// we bought a game
		numGames++
		// reduce s (our money) by game's cost p
		s -= p
		// next game reduces price by d
		p -= d
	}

	// reset p to minimum price of m
	if p < m {
		p = m
	}

	// we can now buy games for p with our remaining money s
	r := s / p

	fmt.Fprintf(writer, "%d\n", numGames+r)
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
