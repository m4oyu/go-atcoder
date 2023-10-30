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
	m, _ := strconv.Atoi(sc.Text())

	var a []int
	for i := 0; i < n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	sort.Ints(a)
	// fmt.Println(a)

	sum := 0
	bottom, top := 0, 0
	for a[top] < a[bottom]+m {
		top++
		sum++
		if n <= top {
			break
		}
	}

	ans := sum
	for i := a[bottom]; top < n; {
		move := a[top] - a[top-1]
		i += move

		top++
		sum++

		for a[bottom] < i {
			bottom++
			sum--
		}

		if ans < sum {
			ans = sum
		}
	}

	fmt.Println(ans)
}
