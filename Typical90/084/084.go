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
	var n int
	var s string
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	s = scanner.Text()

	o := make([]int, n)
	idx_o := 0
	x := make([]int, n)
	idx_x := 0
	for i, v := range s {
		if v == 'o' {
			o[idx_o] = i
			idx_o++
		}
		if v == 'x' {
			x[idx_x] = i
			idx_x++
		}
	}

	if idx_o == 0 || idx_x == 0 {
		fmt.Println(0)
		return
	}

	for i, v := range s {
		// if v == 'o' {

		// }
		if v == 'x' {
			x[idx_x] = i
			idx_x++
		}
	}

	// o, x を配列で管理
	// o[0] = N : 1つ目の o は左からN文字目

}
