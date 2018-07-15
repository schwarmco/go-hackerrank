package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	// your test-case goes here (split by \n)
	in := "8 5\n2 3 1 2 3 2 3 3\n0 3\n4 6\n6 7\n3 5\n0 7"
	out := "1\n2\n3\n2\n1"

	in = "5 5\n1 2 2 2 1\n2 3\n1 4\n2 4\n2 4\n2 3"
	out = "2\n1\n1\n1\n2"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}
