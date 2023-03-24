package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const buf = 200100
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, buf), buf)
}

func debug() {
	inputFile, err := os.Open("./problem.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no inputFile.")
		os.Exit(1)
	}

	scanner = bufio.NewScanner(inputFile)
}

func main() {
	var n, m int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())

	a := make([]int, n+1)
	b := make([]int, m+1)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}
	a[n] = 9223372036854775807
	b[m] = 9223372036854775807

	ai := 0
	bi := 0
	bk := make([]int, m)
	for i := 1; ai != n || bi != m; i++ {
		if a[ai] < b[bi] {
			fmt.Printf("%d ", i)
			ai++
		} else {
			bk[bi] = i
			bi++
		}
	}
	fmt.Println()

	for _, v := range bk {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
