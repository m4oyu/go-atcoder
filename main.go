package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReaderSize(os.Stdin, 1000000)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
}

func debug() {
	input, err := os.Open("./problem.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input.")
		os.Exit(1)
	}
	reader = bufio.NewReader(input)
}

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readint() int {
	ret, err := strconv.Atoi(readLine())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func main() {
	n := readint()
	s := readLine()

	co := -1
	cx := -1
	ans := 0

	for i := 0; i < n; i++ {
		if s[i] == 'o' {
			ans += cx + 1
			co = i
		} else {
			ans += co + 1
			cx = i
		}
	}

	fmt.Println(ans)
}
