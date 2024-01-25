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
	h, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	w, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	board := make([]string, h)
	for i := 0; i < h; i++ {
		sc.Scan()
		board[i] = sc.Text()
	}

	// sumH := make([]int, h)
	// for i := 0; i < h; i++ {
	// 	for
	// }

	ans := math.MaxInt64
	for i := 0; i < h; i++ {
		for j := 0; j < w-k; j++ {
			possible := true
			needToChange := 0
			for t := 0; t < k; t++ {
				target := board[i][j+t]
				if target == 'x' {
					possible = false
					j = j + t + 1
					break
				} else if target == '.' {
					needToChange++
				}
			}

			if possible {
				ans = min(ans, needToChange)
			}
		}
	}

	for j := 0; j < w; j++ {
		for i := 0; i < h-k; i++ {
			possible := true
			needToChange := 0
			for t := 0; t < k; t++ {
				target := board[i+t][j]
				if target == 'x' {
					possible = false
					i = i + t + 1
					break
				} else if target == '.' {
					needToChange++
				}
			}
			if possible {
				ans = min(ans, needToChange)
			}
		}
	}

	if ans == math.MaxInt64 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
