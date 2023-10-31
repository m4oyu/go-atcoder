package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	inputFile, err := os.Open("./in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no inputFile.")
		os.Exit(1)
	}
	sc = bufio.NewScanner(inputFile)
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	var players []Player
	for i := 1; i <= n; i++ {
		sc.Scan()
		r := strings.Split(sc.Text(), "")
		score := 0
		for _, v := range r {
			if v == "o" {
				score++
			}
		}
		players = append(players, Player{id: i, score: score})
	}

	sort.Sort(ByKey(players))

	// slices.SortFunc(players, func(a, b Player) int {
	// 	return cmp.Compare(a.score, b.score)
	// })

	for i := 0; i < len(players); i++ {
		fmt.Printf("%d ", players[i].id)
	}
	fmt.Println()
}

type Player struct {
	id    int
	score int
}

type ByKey []Player

func (s ByKey) Len() int {
	return len(s)
}

func (s ByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
	if s[i].score != s[j].score {
		return s[i].score > s[j].score
	} else {
		return s[i].id < s[j].id
	}
}
