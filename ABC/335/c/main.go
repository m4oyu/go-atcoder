package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

type pair struct {
	x int
	y int
}

func main() {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	queue := make([]pair, 0)
	for i := n; i > 0; i-- {
		queue = append(queue, pair{i, 0})
	}

	for i := 0; i < q; i++ {
		sc.Scan()
		q1 := sc.Text()
		sc.Scan()
		q2 := sc.Text()

		if q1 == "1" {
			switch q2 {
			case "R":
				queue = append(queue, pair{
					x: queue[n-1].x + 1,
					y: queue[n-1].y,
				})
			case "L":
				queue = append(queue, pair{
					x: queue[n-1].x - 1,
					y: queue[n-1].y,
				})
			case "U":
				queue = append(queue, pair{
					x: queue[n-1].x,
					y: queue[n-1].y + 1,
				})
			case "D":
				queue = append(queue, pair{
					x: queue[n-1].x,
					y: queue[n-1].y - 1,
				})
			}

			queue = queue[1:]

		} else {
			index, _ := strconv.Atoi(q2)
			fmt.Printf("%d %d\n", queue[n-index].x, queue[n-index].y)
		}
	}
}
