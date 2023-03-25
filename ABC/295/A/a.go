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

func scan() string {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
	return scanner.Text()
}

func main() {
	n, _ := strconv.Atoi(scan())
	for i := 0; i < n; i++ {
		word := scan()
		for _, v := range []string{"and", "not", "that", "the", "you"} {
			if v == word {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
