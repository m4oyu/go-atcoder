package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

var mod int = int(math.Pow(10, 9)) + 7

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Pow(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

func main() {
	var L, R int
	scanner.Scan()
	L, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	R, _ = strconv.Atoi(scanner.Text())

	var ans int = 0
	for i := 0; i < 18; i++ {
		left := Max(Pow(10, i), L) - 1
		right := Min(Pow(10, i+1)-1, R)

		if left < right {
			var l_ans, r_ans int
			if left%2 == 0 {
				l_ans = (left / 2 % mod) * ((left + 1) % mod) % mod
			} else {
				l_ans = (left % mod) * ((left + 1) / 2 % mod) % mod
			}
			if right%2 == 0 {
				r_ans = (right / 2 % mod) * ((right + 1) % mod) % mod
			} else {
				r_ans = (right % mod) * ((right + 1) / 2 % mod) % mod
			}
			ans = (ans%mod + (((r_ans-l_ans+mod)%mod)*(i+1))%mod) % mod
		}
	}

	RMax := Pow(10, 18)
	if R == RMax {
		ans += ((RMax % mod) * 19) % mod
	}

	fmt.Println(ans % mod)
}
