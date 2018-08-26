package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	// your test-case goes here (split by \n)
	in := "2\n643 953 522 277 464 366 321 409 227 702\n877 654 2 715 229 255 712 267 19 832"
	out := "877 633 491 596 839 875 923 461 27 826\n101 966 573 339 784 718 949 934 62 368"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}
