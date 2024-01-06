package main

import (
	"bufio"
	"fmt"
	"math"
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
	a := make([][]int, 9)
	for i := 0; i < 9; i++ {
		r := make([]int, 9)
		for j := 0; j < 9; j++ {
			sc.Scan()
			r[j], _ = strconv.Atoi(sc.Text())
		}
		a[i] = r
	}

	for i := 0; i < 9; i++ {
		r := make([]int, 9)
		copy(r, a[i])
		sort.Ints(r)
		for j := 0; j < 9; j++ {
			if r[j] != j+1 {
				fmt.Println("No")
				return
			}
		}
	}

	for i := 0; i < 9; i++ {
		r := make([]int, 9)
		for j := 0; j < 9; j++ {
			r[j] = a[j][i]
		}
		sort.Ints(r)
		for j := 0; j < 9; j++ {
			if r[j] != j+1 {
				fmt.Println("No")
				return
			}
		}
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			r := make([]int, 0)
			for x := 3*i - 2; x <= 3*i; x++ {
				for y := 3*i - 2; y <= 3*i; y++ {
					r = append(r, a[x-1][y-1])
				}
			}
			sort.Ints(r)
			for k := 0; k < 9; k++ {
				if r[k] != k+1 {
					fmt.Println("No")
					return
				}
			}

		}
	}

	fmt.Println("Yes")
}
