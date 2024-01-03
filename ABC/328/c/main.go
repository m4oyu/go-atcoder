package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func debug() {
	input, err := os.Open("./input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input file")
		os.Exit(1)
	}
	sc = bufio.NewScanner(input)
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	s := sc.Text()

	sum := make([]int, n+1)
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1]
		if s[i] == s[i-1] {
			sum[i]++
		}
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		l, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		r, _ := strconv.Atoi(sc.Text())

		fmt.Println(sum[r-1] - sum[l-1])
	}

	fmt.Println(bufio.MaxScanTokenSize)
	fmt.Println(math.MaxInt64)
}
