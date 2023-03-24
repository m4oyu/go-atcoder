package main

import (
	"bufio"
	"container/heap"
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

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// var n int
	var q int
	scanner.Scan()
	// n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	q, _ = strconv.Atoi(scanner.Text())

	// 呼ばれた人、行った人を管理する

	minNotCalled := 1 // 1

	recallTarget := &intHeap{}
	heap.Init(recallTarget)
	alreadyGo := &intHeap{}
	heap.Init(alreadyGo)

	for i := 0; i < q; i++ {
		var e, x int
		scanner.Scan()
		e, _ = strconv.Atoi(scanner.Text())

		if e == 1 {
			heap.Push(recallTarget, minNotCalled)
			minNotCalled++
		} else if e == 2 {
			scanner.Scan()
			x, _ = strconv.Atoi(scanner.Text())

			heap.Push(alreadyGo, x)

		} else {
			if alreadyGo.Len() > 0 {
				target := (*recallTarget)[0]
				for target == (*alreadyGo)[0] {
					heap.Pop(recallTarget)
					heap.Pop(alreadyGo)
					target = (*recallTarget)[0]
					if alreadyGo.Len() == 0 {
						break
					}
				}
			}
			fmt.Println((*recallTarget)[0])

		}
	}

}
