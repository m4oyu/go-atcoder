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

func scan() string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner eror: %q\n", err)
	}
	return scanner.Text()
}

func solve(middle int, a []int) int {
	x := 0
	y := 0
	for i := 1; i < len(a); i++ {
		x += a[i] - a[i-1]
		if middle <= x {
			x = 0
			y++
		}
	}
	return y
}

func main() {
	n, _ := strconv.Atoi(scan())
	l, _ := strconv.Atoi(scan())
	k, _ := strconv.Atoi(scan())

	a := make([]int, n+2)
	a[0] = 0
	a[n+1] = l

	for i := 1; i < n+1; i++ {
		a[i], _ = strconv.Atoi(scan())
	}

	left := 1
	right := l/(k+1) + 1

	for left < right {
		middle := (left+right)/2 + 1
		num := solve(middle, a)

		if k+1 <= num {
			left = middle
		} else {
			right = middle - 1
		}
	}

	fmt.Println(right)

}
