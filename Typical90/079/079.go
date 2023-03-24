package main

import (
	"bufio"
	"fmt"
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

func main() {
	var h, w int
	scanner.Scan()
	h, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	w, _ = strconv.Atoi(scanner.Text())
	a := make([][]int, h)
	for i, _ := range a {
		tmp := make([]int, w)
		for j := 0; j < w; j++ {
			scanner.Scan()
			tmp[j], _ = strconv.Atoi(scanner.Text())
		}
		a[i] = tmp
	}
	b := make([][]int, h)
	for i, _ := range a {
		tmp := make([]int, w)
		for j := 0; j < w; j++ {
			scanner.Scan()
			tmp[j], _ = strconv.Atoi(scanner.Text())
		}
		b[i] = tmp
	}

	count := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			diff := a[i][j] - b[i][j]
			if diff != 0 {
				a[i][j] -= diff
				a[i+1][j] -= diff
				a[i][j+1] -= diff
				a[i+1][j+1] -= diff

				if diff < 0 {
					diff = -diff
				}
				count += diff
			}
		}
	}

	for i := 0; i < h; i++ {
		if a[i][w-1] != b[i][w-1] {
			fmt.Println("No")
			return
		}
	}
	for i := 0; i < w; i++ {
		if a[h-1][i] != b[h-1][i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
	fmt.Println(count)
}
