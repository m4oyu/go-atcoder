package main

import (
	"bufio"
	"fmt"
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
	const buf = 200100
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, buf), buf)
}

func debug() {
	inputFile, err := os.Open("./in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no inputFile.")
		os.Exit(1)
	}
	sc = bufio.NewScanner(inputFile)
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())

		if a <= l {
			fmt.Printf("%d ", l)
		} else if l < a && a < r {
			fmt.Printf("%d ", a)
		} else {
			fmt.Printf("%d ", r)
		}
	}
	fmt.Println()
}
