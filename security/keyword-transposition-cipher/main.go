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

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func solve(reader io.Reader, writer io.Writer) {

	N, _ := strconv.Atoi(readLine(reader))

	for N > 0 {

		keyword := readLine(reader)
		ciphertext := readLine(reader)
		_ = ciphertext

		transposition := transpose(keyword)

		var result string
		for _, character := range strings.Split(ciphertext, "") {
			if strings.Contains(transposition, character) {
				key := strings.Index(transposition, character)
				result += string(ALPHABET[key])
			} else {
				result += character
			}
		}

		fmt.Fprintln(writer, result)

		N--
	}
}

func transpose(keyword string) string {

	alphabet := ALPHABET
	keyword = reduceToUnique(keyword)

	// remove keyword-runes from alphabet
	for i := range keyword {
		alphabet = strings.Replace(alphabet, string(keyword[i]), "", -1)
	}

	// build columns
	columns := strings.Split(keyword, "")
	for i := range keyword {
		for j := i; j < len(alphabet); j += len(keyword) {
			columns[i] += string(alphabet[j])
		}
	}

	// sort columns
	sort.Strings(columns)

	return strings.Join(columns, "")
}

func reduceToUnique(keyword string) string {
	e := map[rune]bool{}
	var k string
	for _, v := range keyword {
		if _, ok := e[v]; !ok {
			e[v] = true
			k += string(v)
		}
	}
	return k
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
