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
	input, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sc = bufio.NewScanner(input)
}

var max = int(math.Pow10(18))

func main() {
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())

	for i := 1; i < 20; i++ {
		a := 1
		for j := 1; j <= i; j++ {
			a *= i
		}
		if max < a {
			fmt.Println(-1)
			return
		}
		if a == b {
			fmt.Println(i)
			return
		}
	}
}
