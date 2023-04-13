//lint:file-ignore U1000 Ignore all unused code, it's generated

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReaderSize(os.Stdin, 1000000)
	maxInt = 9223372036854775807
	mod    = 1000000007
)

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "debug" {
			debug()
		}
	}
}

func debug() {
	input, err := os.Open("./problem.in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "There is no input.")
		os.Exit(1)
	}
	reader = bufio.NewReader(input)
}

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readInt() int {
	ret, err := strconv.Atoi(readLine())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}

func readInt2() (int, int) {
	ret := strings.Split(readLine(), " ")
	a, err := strconv.Atoi(ret[0])
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(ret[1])
	if err != nil {
		fmt.Println(err)
	}
	return a, b
}

func readIntArray() []int {
	ret := make([]int, 0)
	tmp := strings.Split(readLine(), " ")
	for _, s := range tmp {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		ret = append(ret, i)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	h, w := readInt2()
	q := readInt()

	color := make([]int, h*w)
	for i := 0; i < h*w; i++ {
		color[i] = 0
	}

	uf := NewUnionFind(h * w)

	diff := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < q; i++ {
		query := readIntArray()
		fmt.Println(query)

		if query[0] == 1 {
			r := query[1]
			c := query[2]
			target := w*(r-1) + c - 1
			color[target] = 1

			for _, v := range diff {
				dr := r + v[0]
				dc := c + v[1]
				if dr <= 0 || h < dr || dc <= 0 || w < dc {
					continue
				}
				if color[w*(dr-1)+dc-1] == 1 {
					uf.unite(target, w*(dr-1)+dc-1)
				}
			}
		}

		if query[0] == 2 {
			ra := query[1]
			ca := query[2]
			rb := query[3]
			cb := query[4]

			if uf.same(w*(ra-1)+ca-1, w*(rb-1)+cb-1) && color[w*(ra-1)+ca-1] == 1 && color[w*(rb-1)+cb-1] == 1 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

	}

}

func NewUnionFind(n int) *Union_Find {
	tree := make([]int, n)
	for i := 0; i < n; i++ {
		tree[i] = i
	}
	return &Union_Find{
		par: tree,
	}
}

type Union_Find struct {
	par []int
}

func (uf *Union_Find) root(x int) int {
	if uf.par[x] == x {
		return x
	}

	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]

}

func (uf *Union_Find) unite(x, y int) {
	rx := uf.root(x)
	ry := uf.root(y)
	if rx == ry {
		return
	}
	uf.par[rx] = ry
}

func (uf *Union_Find) same(x, y int) bool {
	rx := uf.root(x)
	ry := uf.root(y)
	return rx == ry
}
