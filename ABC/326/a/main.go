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
	x, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	y, _ := strconv.Atoi(sc.Text())

	diff := x - y
	switch diff {
	case -2, -1, 1, 2, 3:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
	}

}
