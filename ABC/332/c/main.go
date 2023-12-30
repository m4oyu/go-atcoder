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
	sc.Text()
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s := sc.Text()
	s += "0"

	sum := m
	logo := 0

	i := 0
	j := 0
	for _, v := range s {
		switch v {
		case '0':
			if sum < i {
				sum = i
			}
			if logo < j {
				logo = j
			}
			i = 0
			j = 0
		case '1':
			i++
		case '2':
			i++
			j++
		}
	}

	if sum-m < logo {
		fmt.Println(logo)
	} else {
		fmt.Println(sum - m)
	}

}
