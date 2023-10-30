package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

var h, w int
var s [][]int
var check [][]int
var q = make([]int, 0)
var dh = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dw = []int{-1, 0, 1, -1, 1, -1, 0, 1}
var ans int = 0

func main() {
	sc.Scan()
	h, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	w, _ = strconv.Atoi(sc.Text())

	for i := 0; i < h; i++ {
		sc.Scan()
		tmp := strings.Split(sc.Text(), "")
		var add []int
		for _, v := range tmp {
			if v == "#" {
				add = append(add, 1)
			} else {
				add = append(add, 0)
			}
		}
		s = append(s, add)
	}

	for i := 0; i < h; i++ {
		check = append(check, make([]int, w))
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if check[i][j] == 0 && s[i][j] == 1 {
				q = append(q, i*w+j)
				check[i][j] = 1
				ans++
			}
			for 0 < len(q) {
				solve(q[0])
				q = q[1:]
			}
		}
	}
	fmt.Println(ans)
}

func solve(v int) {
	height := v / w
	width := v % w

	for i := 0; i < 8; i++ {
		if height+dh[i] < 0 || h <= height+dh[i] || width+dw[i] < 0 || w <= width+dw[i] {
			continue
		}
		if check[height+dh[i]][width+dw[i]] == 0 && s[height+dh[i]][width+dw[i]] == 1 {
			q = append(q, (height+dh[i])*w+(width+dw[i]))
			check[height+dh[i]][width+dw[i]] = 1
		}
	}
}
