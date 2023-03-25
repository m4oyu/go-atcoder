package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
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

func scan() string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner eror: %q\n", err)
	}
	return scanner.Text()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var r, c int
var area [][]rune

func bomb(y int, x int, l int) {
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if Abs(i-y)+Abs(j-x) <= l && area[i][j] == '#' {
				area[i][j] = '.'
			}
		}
	}
}

func main() {
	r, _ = strconv.Atoi(scan())
	c, _ = strconv.Atoi(scan())

	area = make([][]rune, r)
	for i := 0; i < r; i++ {
		s := scan()
		col := make([]rune, c)
		for j := 0; j < c; j++ {
			a, b := utf8.DecodeRuneInString(s)
			col[j] = a
			s = s[b:]
		}
		area[i] = col
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if area[i][j] != '#' && area[i][j] != '.' {
				bomb(i, j, int(area[i][j]-'0'))
				area[i][j] = '.'
			}
		}
	}

	for i := 0; i < r; i++ {
		fmt.Println(string(area[i]))
	}
}
