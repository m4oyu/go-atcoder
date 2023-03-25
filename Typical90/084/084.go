package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode/utf8"
)

var scanner = bufio.NewScanner(os.Stdin)
var reader = bufio.NewReaderSize(os.Stdin, 4*1024)

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
	input, err := os.Open("./problem.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input.")
		os.Exit(1)
	}

	scanner = bufio.NewScanner(input)
	reader = bufio.NewReader(input)
	defer input.Close()
}

func scan() string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
	return scanner.Text()
}

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := reader.ReadLine()
		if e == io.EOF {
			break
		}
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	var n int
	var s string
	n, _ = strconv.Atoi(scan())
	s = readLine()

	var a []int
	tmp, _ := utf8.DecodeRuneInString(s[0:])
	count := 0
	for _, v := range s {
		if tmp == v {
			count++
		} else {
			a = append(a, count)
			tmp = v
			count = 1
		}
	}
	a = append(a, count)

	var extra int = 0
	for _, v := range a {
		extra += v * (v + 1) / 2
	}

	fmt.Println((n * (n + 1) / 2) - extra)

}
