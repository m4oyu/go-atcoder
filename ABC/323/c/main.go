package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
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
	n := scanInt()
	m := scanInt()

	a := make([]int, m)
	for i := 0; i < m; i++ {
		a[i] = scanInt()
	}

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = scanString()
	}

	type result struct {
		score    int
		solved   []int
		unsolved []int
	}

	results := make([]result, n)
	max := -1

	for i := 0; i < n; i++ {
		solved := make([]int, 0)
		unsolved := make([]int, 0)
		score := i + 1
		for index, v := range s[i] {
			if v == 'o' {
				solved = append(solved, a[index])
				score += a[index]
			} else {
				unsolved = append(unsolved, a[index])
			}
		}

		if max < score {
			max = score
		}

		sort.Ints(solved)
		sort.Slice(unsolved, func(i, j int) bool {
			return unsolved[i] > unsolved[j]
		})
		results[i] = result{score: score, solved: solved, unsolved: unsolved}
	}

	for i := 0; i < n; i++ {
		ans := 0
		score := results[i].score
		if max == score {
			fmt.Println(ans)
			continue
		}
		for j := 0; j < len(results[i].unsolved); j++ {
			score += results[i].unsolved[j]
			ans++
			if max <= score {
				fmt.Println(ans)
				break
			}
		}
	}
}

func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func scanString() string {
	sc.Scan()
	return sc.Text()
}
