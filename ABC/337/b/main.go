package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func debug() {
	f, err := os.Open("./input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sc = bufio.NewScanner(f)
}

func main() {
	sc.Scan()
	s := sc.Text()

	abc := byte(s[0])

	for i := 0; i < len(s); i++ {
		if abc != s[i] {
			if abc == 'A' && s[i] == 'B' {
				abc = 'B'
			} else if abc == 'A' && s[i] == 'C' {
				abc = 'C'
			} else if abc == 'B' && s[i] == 'C' {
				abc = 'C'
			} else {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
}
