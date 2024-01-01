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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	sorta := make([]int, n+1)
	copy(sorta, a)
	sort.Slice(sorta, func(i, j int) bool {
		return sorta[i] > sorta[j]
	})

	sum := make([]int, 1000009)
	tmp := 0
	for i := 0; i < n; i++ {
		tmp += sorta[i]
		for j := sorta[i]; sorta[i+1] <= j; j-- {
			sum[j] = tmp
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", sum[a[i]+1])
	}
	fmt.Println()

}
