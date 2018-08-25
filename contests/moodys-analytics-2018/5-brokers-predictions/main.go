package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type correlation []int

func (c correlation) contains(i int) bool {
	for _, a := range c {
		if a == i {
			return true
		}
	}
	return false
}

func solve(reader io.Reader, writer io.Writer) {

	Predictions := map[int][]correlation{
		0: []correlation{},
	}

	nq := strings.Split(readLine(reader), " ")
	// n, _ := strconv.Atoi(nq[0])
	q, _ := strconv.Atoi(nq[1])

	for t := 1; t <= q; t++ {

		line := strings.Split(readLine(reader), " ")

		if line[0] == "1" { // prediction

			pt, _ := strconv.Atoi(line[1])
			si, _ := strconv.Atoi(line[2])
			sj, _ := strconv.Atoi(line[3])

			// find correlationsets containing si or sj
			found := []int{}
			for idx, set := range Predictions[pt] {
				if set.contains(si) || set.contains(sj) {
					found = append(found, idx)
					if len(found) > 1 {
						break
					}
				}
			}

			if len(found) > 1 {
				// we need to merge 2 sets
				Predictions[t] = make([]correlation, len(Predictions[pt]))
				copy(Predictions[t], Predictions[pt])
				Predictions[t][found[0]] = append(Predictions[t][found[0]], Predictions[t][found[1]]...)
				Predictions[t] = append(Predictions[t][:found[1]], Predictions[t][found[1]+1:]...)
			} else if len(found) > 0 {
				// we can just add it to the 1 found correlation set
				idx := found[0]
				Predictions[t] = make([]correlation, len(Predictions[pt]))
				copy(Predictions[t], Predictions[pt])
				Predictions[t][idx] = append(Predictions[t][idx], si, sj)
			} else {
				// we need to create a new correlation set
				Predictions[t] = append(Predictions[pt], correlation{si, sj})
			}

		} else { // request

			pt, _ := strconv.Atoi(line[1])
			kt, _ := strconv.Atoi(line[2])

			var count int
			groups := make([]int, len(Predictions[pt]))

			S := map[int]bool{}
			for i := 0; i < kt; i++ {
				idx, _ := strconv.Atoi(line[i+3])
				S[idx] = false
				for p, set := range Predictions[pt] {
					if set.contains(idx) {
						groups[p]++
						S[idx] = true
						break
					}
				}
			}

			// non-tagged stocks can moved independently
			for s := range S {
				if !S[s] {
					count++
				}
			}

			// add correlationsets, we hit
			for _, g := range groups {
				if g > 0 {
					count++
				}
			}

			numScenarios := int((math.Pow(2, float64(count))))
			outputFilter := int(math.Pow(10, 9) + 7)

			fmt.Fprintln(writer, numScenarios%outputFilter)
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
