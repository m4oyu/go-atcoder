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

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		r := make([]int, n)
		m[i] = r
	}
	m[(n-1)/2][(n-1)/2] = -1
	m[0][0] = 1

	x := 0
	y := 0
	dir := "R"
	for i := 2; i < n*n; i++ {
		switch dir {
		case "R":
			if x+1 > n-1 {
				dir = "D"
				y = y + 1
			} else if m[y][x+1] != 0 {
				dir = "D"
				y = y + 1
			} else {
				x = x + 1
			}
		case "L":
			if x-1 < 0 {
				dir = "U"
				y = y - 1
			} else if m[y][x-1] != 0 {
				dir = "U"
				y = y - 1
			} else {
				x = x - 1
			}
		case "U":
			if y-1 < 0 {
				dir = "R"
				x = x + 1
			} else if m[y-1][x] != 0 {
				if m[y][x+1] == -1 {
					fmt.Println("end")
				} else {
					dir = "R"
					x = x + 1
				}
			} else {
				y = y - 1
			}

		case "D":
			if y+1 > n-1 {
				dir = "L"
				x = x - 1
			} else if m[y+1][x] != 0 {
				dir = "L"
				x = x - 1
			} else {
				y = y + 1
			}
		}
		m[y][x] = i

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == -1 {
				fmt.Printf("%s ", "T")
			} else {
				fmt.Printf("%d ", m[i][j])
			}
		}
		fmt.Println()
	}

}
