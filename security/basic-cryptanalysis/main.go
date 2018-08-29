package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"
const PLACEHOLDER = "_"

type transposition struct {
	Characters []byte
}

func NewTransposition() transposition {
	return transposition{
		[]byte(strings.Repeat(PLACEHOLDER, len(ALPHABET))),
	}
}

func (t transposition) decode(word string) (string, error) {

	// if len(t.From) != len(ALPHABET) || len(t.To) != len(ALPHABET) {
	// 	return "", errors.New("incomplete transposition")
	// }

	var result []byte
	for _, char := range word {
		pos := strings.IndexRune(ALPHABET, char)
		result = append(result, t.Characters[pos])
	}

	return string(result), nil
}

func (t transposition) String() string {
	return string(t.Characters)
}

func (t transposition) extend(from, to string) error {

	for i := range from {
		pos := strings.IndexByte(ALPHABET, from[i])
		if string(t.Characters[pos]) != PLACEHOLDER && t.Characters[pos] != to[i] {
			return fmt.Errorf("confilct in transposition: want(%s > %s), already have(%s > %s)",
				string(from[i]),
				string(to[i]),
				string(from[i]),
				string(t.Characters[i]))
		}
		t.Characters[pos] = to[i]
	}

	return nil
}

func (t transposition) merge(m transposition, force bool) (transposition, error) {
	trans := NewTransposition()
	for i := range t.Characters {
		// merge conflict?
		if !force && string(m.Characters[i]) != PLACEHOLDER && string(t.Characters[i]) != PLACEHOLDER && t.Characters[i] != m.Characters[i] {
			return NewTransposition(), fmt.Errorf("conflict in merging %v and %v", t, m)
		}
		// m is first, because it may be forced
		if string(m.Characters[i]) != PLACEHOLDER {
			trans.Characters[i] = m.Characters[i]
		} else if string(t.Characters[i]) != PLACEHOLDER {
			trans.Characters[i] = t.Characters[i]
		} else {
			trans.Characters[i] = []byte(PLACEHOLDER)[0]
		}
	}
	return trans, nil
}

func (t transposition) validate(ciphers []string, dict map[int][]string) bool {

	for _, cipher := range ciphers {

		decoded, _ := t.decode(cipher)

		// skip, if nothing was transposed (got placeholders only)
		if decoded == strings.Repeat(PLACEHOLDER, len(cipher)) {
			continue
		}

		var found bool

		// check, if decoded word is found in dictionary
		r := regexp.MustCompile("^" + strings.Replace(decoded, PLACEHOLDER, ".", -1) + "$")
		for _, word := range dict[len(cipher)] {
			if r.MatchString(word) {
				found = true
				break
			}
		}

		// nothing found? this transposition can't be valid
		if !found {
			return false
		}

	}

	// we found a possible dict-word for every decoded cipher
	return true
}

func solve(reader io.Reader, writer io.Writer) {

	file, _ := ioutil.ReadFile("dictionary.lst")
	dict := strings.Split(strings.ToLower(string(file)), "\n")

	// read dict into a map sorted by word-length
	dictByLength := make(map[int][]string)
	for _, d := range dict {
		dictByLength[len(d)] = append(dictByLength[len(d)], d)
	}

	// read ciphers from stdin
	ciphers := strings.Split(readLine(reader), " ")

	validTranspositions := make(map[string]transposition)

	// iterate over all ciphers
	for _, cipher := range ciphers {

		// iterate over the dictionary (only same-length words)
		for _, word := range dictByLength[len(cipher)] {

			// skip, if not isomorphic
			if !isIsomorphic(cipher, word) {
				continue
			}

			// create a partial transposition for word -> cipher
			trans := NewTransposition()
			trans.extend(cipher, word)

			// skip, if we already have this transposition
			if _, ok := validTranspositions[trans.String()]; ok {
				continue
			}

			// test it against all (same-length) dict-words
			if trans.validate(ciphers, dictByLength) {
				validTranspositions[trans.String()] = trans
			}

		}
	}

	mergedTranspositions := make(map[string]transposition)

	dummy := NewTransposition()
	mergedTranspositions[string(dummy.Characters)] = dummy

	// we now have to merge all valid partial transpositions
	for _, part := range validTranspositions {

		for str, t := range mergedTranspositions {

			merged, err := t.merge(part, false)

			if err == nil {
				if merged.String() == str {
					continue
				}
				if merged.validate(ciphers, dictByLength) {
					// add new (merged) transposition and delete old one
					mergedTranspositions[string(merged.Characters)] = merged
					delete(mergedTranspositions, str)
					continue
				}
			} else {
				// fork transposition
				merged, _ = t.merge(part, true)
				if merged.String() != str {
					if merged.validate(ciphers, dictByLength) {
						mergedTranspositions[merged.String()] = merged
						delete(mergedTranspositions, str)
					}
				}
				fork, _ := part.merge(t, true)
				if fork.String() != str {
					if fork.validate(ciphers, dictByLength) {
						mergedTranspositions[fork.String()] = fork
						delete(mergedTranspositions, str)
					}
				}
			}
		}
	}

	var finalTransposition transposition
	for _, t := range mergedTranspositions {
		finalTransposition = t
		break
	}

	// transpose output
	output := []string{}
	for _, cipher := range ciphers {
		word, _ := finalTransposition.decode(cipher)
		output = append(output, word)
	}

	fmt.Fprintln(writer, strings.Join(output, " "))
}

func isIsomorphic(a, b string) bool {

	M := make(map[byte]byte)
	C := make(map[byte]bool)

	for i := range a {
		if _, ok := M[a[i]]; !ok {
			if _, ok := C[b[i]]; ok {
				return false
			}
			C[b[i]] = true
			M[a[i]] = b[i]
		} else if M[a[i]] != b[i] {
			return false
		}
	}

	return true
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
