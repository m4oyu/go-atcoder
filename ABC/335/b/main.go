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
	for i := 0; i <= 21; i++ {
		for j := 0; j <= 21; j++ {
			for k := 0; k <= 21; k++ {
				if i+j+k > n {
					continue
				}
				fmt.Printf("%d %d %d\n", i, j, k)
			}
		}
	}
}
