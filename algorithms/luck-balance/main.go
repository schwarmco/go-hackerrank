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

type contest struct {
	Luck      int
	Important bool
}

func solve(reader io.Reader, writer io.Writer) {

	nk := strings.Split(readLine(reader), " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	_ = k

	var contests []contest
	for n > 0 {

		lt := strings.Split(readLine(reader), " ")
		l, _ := strconv.Atoi(lt[0])
		t, _ := strconv.ParseBool(lt[1])

		contests = append(contests, contest{l, t})

		n--
	}

	sort.Slice(contests, func(i, j int) bool {
		return contests[i].Luck > contests[j].Luck
	})

	var luck int
	for _, c := range contests {

		if c.Important {
			if k <= 0 {
				luck -= c.Luck
				continue
			}
			k--
		}

		luck += c.Luck
	}

	fmt.Fprintln(writer, luck)
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
