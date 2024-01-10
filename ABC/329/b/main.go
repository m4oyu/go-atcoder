package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			input, err := os.Open("./input")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			sc = bufio.NewScanner(input)
		}
	}
	const buf = 200100
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, buf), buf)
}

func main() {
	var repunit []int
	repunit = append(repunit, 1)
	a := "1"
	for i := 0; i < 11; i++ {
		a = a + "1"
		b, _ := strconv.Atoi(a)
		repunit = append(repunit, b)
	}

	var list []int
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			for k := 0; k < 12; k++ {
				sum := repunit[i] + repunit[j] + repunit[k]
				if !contains(list, sum) {
					list = append(list, sum)
				}
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println(list[n-1])

}
func contains(arr []int, val int) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}
