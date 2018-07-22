package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	in := "3\n2147483647\n1\n0"
	out := "2147483648\n4294967294\n4294967295"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}
