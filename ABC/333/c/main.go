package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	const buf = 200100
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, buf), buf)
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

	r := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		r[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	sum := make([]int, n)
	sum[0] = r[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + r[i]
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		query, _ := strconv.Atoi(sc.Text())

		left := -1
		right := n
		middle := (right + left) / 2

		for right-left > 1 {
			if sum[middle] < query {
				left = middle
			} else if sum[middle] == query {
				left = middle
				right = middle + 1
			} else {
				right = middle
			}
			middle = left + ((right - left + 1) / 2)
		}
		fmt.Println(left + 1)
	}

}
