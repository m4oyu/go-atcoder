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

	for i := n; i <= 919; i++ {
		hundred := i / 100
		ten := i % 100 / 10
		one := i % 10

		if hundred*ten == one {
			fmt.Println(i)
			return
		}
	}
}
