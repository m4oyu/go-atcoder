package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h)
	fmt.Scan(&w)

	sc := bufio.NewScanner(os.Stdin)

	var s [][]int
	for i := 0; i < h; i++ {
		sc.Scan()
		tmp := strings.Split(sc.Text(), "")
		var add []int
		for _, v := range tmp {
			if "#" == v {
				add = append(add, 1)
			} else {
				add = append(add, 0)
			}
		}
		s = append(s, add)
	}
	fmt.Println(s)

	m := make(map[int][]int)

	fmt.Println(m)

	dh := []int{-1, -1, -1, 0, 0, 0, 1, 1, 1}
	dw := []int{-1, 0, 1, -1, 0, 1, -1, 0, 1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 1 {
				count := 0
				for k := 0; k < 9; k++ {
					kh := i + dh[k]
					kw := j + dw[k]
					if kh < 0 || h <= kh || kw < 0 || w <= kw {
						break
					}

					_, ok := m[kh*w+kw]
					if s[kh][kw] == 1 && ok {
						count++
					}
				}

				if count == 0 {

				}
			}

		}

	}
}
