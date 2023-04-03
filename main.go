//lint:file-ignore U1000 Ignore all unused code, it's generated

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReaderSize(os.Stdin, 1000000)
var maxInt = 9223372036854775807

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

func readInt() int {
	ret, err := strconv.Atoi(readLine())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func readIntArray() []int {
	ret := make([]int, 0)
	tmp := strings.Split(readLine(), " ")
	for _, s := range tmp {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		ret = append(ret, i)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
}
