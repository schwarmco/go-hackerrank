package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {

	// your test-case goes here (split by \n)
	in := "2\nSPORT\nLDXTW KXDTL NBSFX BFOII LNBHG ODDWN BWK\nSECRET\nJHQSU XFXBQ"
	out := "ILOVE SOLVI NGPRO GRAMM INGCH ALLEN GES\nCRYPT OLOGY"

	input := bufio.NewReader(strings.NewReader(in))
	output := new(bytes.Buffer)

	solve(input, output)

	if out != strings.Trim(output.String(), "\r\n") {
		t.Fatal("\nexpected:\n", out, "\n\ngot:\n", output)
	}
}
