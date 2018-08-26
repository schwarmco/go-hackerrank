package main

import (
	"math/bits"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const n = 1000

func solve(reader io.Reader, writer io.Writer) {

	N, _ := strconv.Atoi(readLine(reader))

	for N > 0 {

		// read first 10 numbers
		a := make([]int, 10)
		for idx, s := range strings.Split(readLine(reader), " ") {
			a[idx], _ = strconv.Atoi(s)
		}

		// the java-rng-implementation leaks the lower bits of the seed into
		// the resulting numbers (via bits % n). We can therefore find the
		// first few bits of seed and bruteforce the rest

		var lowerBits, higherBits, seed uint64

		// find the number of bits which are leaked
		var numLeakedBits int
		numLeakedBits = bits.TrailingZeros(n)

		// find the lower seed part, which matches the a-series
		for lowerBits = uint64(0); lowerBits < uint64(1)<<20; lowerBits++ {
			if isCorrectLowerBits(lowerBits, a, numLeakedBits) {
				// we found the matching lower (right-most) 20 bits! -> lowerBits
				break
			}
		}

		// now we brute-force the remainig 24 high bits (combined with the found lowerBits)
		for higherBits = uint64(1); higherBits < uint64(1)<<24; higherBits++ {
			seed = bits.RotateLeft64(higherBits, 20) | lowerBits
			if isCorrectSeed(seed, a) {
				// we found the correct seed! :)
				break
			}
		}

		// now we need to (re-)generate the first sequence ...
		r := NewRandom(seed)
		for j := 0; j < 10; j ++ {
			r.NextInt(1000)
		}

		// ... to get the following 10 numbers
		var nextSeq []string
		for j := 0; j < 10; j++ {
			nextSeq = append(nextSeq, strconv.Itoa(r.NextInt(1000)))
		}

		fmt.Fprintln(writer, strings.Join(nextSeq, " "))

		N--
	}
}

func isCorrectLowerBits(seed uint64, a []int, len int) bool {
	r := NewRandom(seed)
	
	for i := range a {
		x := r.NextInt(1000)
		for j := 0; j < len; j++ {
			m := int(math.Pow(2, float64(j)))
			if (a[i] & m) != (x & m) {
				return false
			}
		}
	}
	return true
}

func isCorrectSeed(seed uint64, a []int) bool {

	r := NewRandom(seed)

	for depth, i := range a {
		x := r.NextInt(1000)
		if x != i {
			if depth > 2 {
				fmt.Println("failed after", depth, "- got", x, "wanted", i)
			}
			return false
		}
	}

	return true
}

type Random struct {
	Seed uint64
}

func NewRandom(seed uint64) *Random {
	s := (seed ^ 0x5DEECE66D) & ((1 << 48) - 1)
	return &Random{
		Seed: s,
	}
}

func (r *Random) NextInt(n int) int {
	var bits, val uint32
	for {
		bits = r.next(31)
		val = bits % uint32(n)
		if bits-val+(uint32(n)-1) >= 0 {
			break
		}
	}

	return int(val)
}

func (r *Random) next(bits int) uint32 {
	r.Seed = (r.Seed*0x5DEECE66D + 0xB) & ((1 << 48) - 1)
	return uint32(r.Seed >> uint32(48-bits))
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
