package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	// your test-case goes here (split by \n)
	in := "4 5\n1 0 1 2\n2 0 2 1 2\n2 1 2 2 1\n1 1 2 3\n2 4 3 1 2 3"
	out := "4\n2\n2"

	in = "3 6\n1 0 1 2\n1 0 2 3\n1 0 1 3\n2 1 2 1 2\n2 2 2 1 2\n2 3 2 1 2"
	out = "2\n4\n4"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}

func Benchmark_solve(b *testing.B) {

	in := "4 5\n1 0 1 2\n2 0 2 1 2\n2 1 2 2 1\n1 1 2 3\n2 4 3 1 2 3"
	// out := "4\n2\n2"

	for n := 0; n < b.N; n++ {
		input := bufio.NewReader(strings.NewReader(in))
		output := new(bytes.Buffer)
		solve(input, output)
	}
}
