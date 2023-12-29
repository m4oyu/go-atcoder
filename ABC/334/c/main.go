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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ := strconv.Atoi(sc.Text())

	var a []int
	for i := 0; i < k; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}

	ans := 0
	if (2*n-k)%2 == 0 {
		// 偶数
		for i := 0; i < k; i += 2 {
			ans += a[i+1] - a[i]
		}
	} else {
		// 奇数 どれかはじくものを選ぶ
		front := make([]int, k)
		back := make([]int, k)
		front[0] = 0
		back[k-1] = 0
		for i := 2; i < k; i += 2 {
			front[i] = front[i-2] + a[i-1] - a[i-2]
		}
		for i := k - 3; 0 <= i; i -= 2 {
			back[i] = back[i+2] + a[i+2] - a[i+1]
		}

		ans = math.MaxInt
		for i := 0; i < k; i += 2 {
			if ans > front[i]+back[i] {
				ans = front[i] + back[i]
			}
		}
	}
	fmt.Println(ans)
}
