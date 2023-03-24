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
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintln(os.Stderr, "atoi error")
	}

	a := make([]int, n)
	var sum int = 0
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "atoi error")
		}
		sum += a[i]
	}
	if sum%10 != 0 {
		fmt.Println("No")
		return
	}

	a = append(a, a...)
	b := make([]int, 2*n+1)
	b[0] = 0
	for i := 0; i < 2*n; i++ {
		b[i+1] = b[i] + a[i]
	}

	for i := 0; i < n; i++ {
		var left, right int = i + 1, 2*n - 1
		for 0 < right-left {
			middle := (left + right) / 2
			size := b[middle] - b[i]

			if size < sum/10 {
				left = middle + 1
			}
			if size == sum/10 {
				fmt.Println("Yes")
				return
			}
			if size > sum/10 {
				right = middle
			}
		}
	}

	fmt.Println("No")
}
