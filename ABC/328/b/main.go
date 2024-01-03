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

	d := make([]int, 101)
	for i := 1; i <= n; i++ {
		sc.Scan()
		d[i], _ = strconv.Atoi(sc.Text())
	}

	ans := 0
	for i := 1; i <= n; i++ {
		if i < 10 {
			m := i
			for j := 1; j <= d[i]; j++ {
				if j%10 == m && (j/10 == 0 || j/10 == m) {
					ans++
				}
			}
		} else if i == 100 {
			continue
		} else {
			if i/10 != i%10 {
				continue
			}

			m := i / 10
			for j := 1; j <= d[i]; j++ {
				if j%10 == m && (j/10 == 0 || j/10 == m) {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
