package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	node []int
}

func main() {
	var n int
	fmt.Scan(&n)

	var w []int
	var x []int

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		l := strings.Split(sc.Text(), " ")
		wi, _ := strconv.Atoi(l[0])
		xi, _ := strconv.Atoi(l[1])

		w = append(w, wi)
		x = append(x, xi)
	}

	// fmt.Println(w)
	// fmt.Println(x)

	maxAns := 0
	for i := 0; i < 24; i++ {
		ans := 0
		for j := 0; j < n; j++ {
			t := (i + x[j]) % 24
			if 9 <= t && t < 18 {
				ans += w[j]
			}
		}
		if maxAns < ans {
			maxAns = ans
		}
	}
	fmt.Println(maxAns)
}
