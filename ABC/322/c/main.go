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
	f, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sc = bufio.NewScanner(f)
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i := 0; i < m; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		ans[a-1] = 0
	}

	count := 0
	for i := n - 1; i >= 0; i-- {
		if ans[i] == 0 {
			count = 0
		} else {
			count++
			ans[i] = count
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}
