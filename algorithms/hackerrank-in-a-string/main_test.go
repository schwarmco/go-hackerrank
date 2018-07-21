package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
    
    // your test-case goes here (split by \n)
	in := "2\nhereiamstackerrank\nhackerworld"
	out := "YES\nNO"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}
