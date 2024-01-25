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

	a := make(map[int]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		preHuman, _ := strconv.Atoi(sc.Text())
		a[preHuman] = i + 1
	}

	target := -1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[target])
		target = a[target]
	}
	fmt.Println()
}
