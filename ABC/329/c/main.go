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
	sc.Scan()
	s := sc.Text()

	m := make(map[byte]int)
	tmp := byte('1')
	count := 1

	for i := 0; i < n; i++ {
		if tmp == s[i] {
			count++
		} else {
			count = 1
		}

		tmp = s[i]

		if _, ok := m[tmp]; !ok {
			m[tmp] = 1
		} else {
			if m[tmp] < count {
				m[tmp] = count
			}
		}
	}

	ans := 0
	for _, v := range m {
		ans += v
	}
	fmt.Println(ans)
}
