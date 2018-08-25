package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	// your test-case goes here (split by \n)
	in := "2\n2 2\n3 2"
	out := "1\n2"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}

func Benchmark_solve(b *testing.B) {

	in := "2\n2 2\n3 2"

	for n := 0; n < b.N; n++ {
		input := bufio.NewReader(strings.NewReader(in))
		output := new(bytes.Buffer)
		solve(input, output)
	}
}
