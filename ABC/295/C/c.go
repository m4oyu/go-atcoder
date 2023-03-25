package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func main() {
	n, _ := strconv.Atoi(scan())

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(scan())
	}

	sort.Ints(a)

	ans := 0
	color := a[0]
	count := 1

	for i := 1; i < n; i++ {
		if color == a[i] {
			count++
		} else {
			ans += count / 2
			color = a[i]
			count = 1
		}
	}
	ans += count / 2

	fmt.Println(ans)

}
