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
	input, err := os.Open("./input")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input file")
		os.Exit(1)
	}
	sc = bufio.NewScanner(input)
}

func main() {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())

	ans := 0
	if r < a {
		ans = ((r - l) + ((a - r) % m)) / m
		if (a-r)%m == 0 {
			ans++
		}
	} else if l <= a && a <= r {
		ans += ((r - a) / m) + ((a - l) / m) + 1
	} else {
		ans = ((r - l) + ((l - a) % m)) / m
		if (l-a)%m == 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
