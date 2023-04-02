package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	BUFSIZE = 10000000
	MOD     = 1000000007
)

var rdr *bufio.Reader

func main() {
	rdr = bufio.NewReaderSize(os.Stdin, BUFSIZE)
	solve()
}

func solve() {
	n := readint()
	s := readline()
	a := make([]int, 0)
	m := make(map[byte]int)
	cnt := 1
	for i := 0; i < n-1; i++ {
		m[s[i]]++
		if s[i] == s[i+1] {
			cnt++
		} else {
			a = append(a, (cnt*cnt+cnt)/2)
			cnt = 1
		}
	}
	m[s[n-1]]++
	if cnt != 0 {
		a = append(a, (cnt*cnt+cnt)/2)
	}
	if len(m) != 2 {
		fmt.Println(0)
		return
	}
	ans := (n*n + n) / 2

	for i := 0; i < len(a); i++ {
		ans -= a[i]
	}
	fmt.Println(ans)
}

// 汎用関数
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(x ...int) int {
	ret := 0
	for i := 0; i < len(x); i++ {
		if ret < x[i] {
			ret = x[i]
		}
	}
	return ret
}
func min(x ...int) int {
	ret := math.MaxInt64
	for i := 0; i < len(x); i++ {
		if ret > x[i] {
			ret = x[i]
		}
	}
	return ret
}
func printList(l []int) {
	s := make([]string, 0)
	for i := 0; i < len(l); i++ {
		s = append(s, strconv.Itoa(l[i]))
	}
	fmt.Println(strings.Join(s, " "))
}
func sumlist(l []int) int {
	sum := 0
	for i := 0; i < len(l); i++ {
		sum += l[i]
	}
	return sum
}
func countGroup(a []int, x int) int {
	cnt := 0
	t := 0
	for i := 0; i < len(a); i++ {
		if a[i]-x >= 0 {
			t++
			continue
		} else if t != 0 {
			cnt++
			t = 0
		}
	}
	if t != 0 {
		cnt++
	}
	return cnt
}

// 標準入出力
func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return strings.TrimSpace(string(buf))
}
func readIntSliceBanpei() []int {
	slice := make([]int, 0)
	slice = append(slice, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	slice = append(slice, math.MaxInt64)
	return slice
}
func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2i(v))
	}
	return slice
}
func readFloatSlice() []float64 {
	slice := make([]float64, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, s2f(v))
	}
	return slice
}
func readint() int {
	return s2i(readline())
}
func readfloat() float64 {
	return s2f(readline())
}
func readint2() (int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1])
}
func readfloat2() (float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1])
}
func readint3() (int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2])
}
func readfloat3() (float64, float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1]), s2f(lines[2])
}
func readint4() (int, int, int, int) {
	lines := strings.Split(readline(), " ")
	return s2i(lines[0]), s2i(lines[1]), s2i(lines[2]), s2i(lines[3])
}
func readfloat4() (float64, float64, float64, float64) {
	lines := strings.Split(readline(), " ")
	return s2f(lines[0]), s2f(lines[1]), s2f(lines[2]), s2f(lines[3])
}
func s2i(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic("Faild : " + s + " can't convert to int")
	}
	return v
}
func s2f(s string) float64 {
	v, ok := strconv.ParseFloat(s, 64)
	if ok != nil {
		panic("Faild : " + s + " can't convert to float")
	}
	return v
}
func readSimpleUndirectedGraph(n, m int) [][]int {
	g := make([][]int, n+1)
	for i := 0; i < m; i++ {
		a, b := readint2()
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	return g
}

// 2分探索
func UpperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
func LowerBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] >= target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// 変換
func stob(s string) []byte {
	ret := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ret[i] = s[i]
	}
	return ret
}
func mtol(m map[int]int) [][]int {
	t := make([][]int, len(m))
	idx := 0
	for i, v := range m {
		t[idx] = []int{i, v}
		idx++
	}
	sort.Slice(t, func(i, j int) bool { return t[i][0] < t[j][0] })
	return t
}
func ltom(l []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(l); i++ {
		m[l[i]]++
	}
	return m
}
func reverseSort(a []int) []int {
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	return a
}
func reverseList(x []int) []int {
	ret := make([]int, 0)
	for i := len(x) - 1; i >= 0; i-- {
		ret = append(ret, x[i])
	}
	return ret
}
func intSlice3(a, b, c int) [][][]int {
	ret := make([][][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([][]int, b)
		for j := 0; j < b; j++ {
			ret[i][j] = make([]int, c)
		}
	}
	return ret
}
func intSlice2(a, b int) [][]int {
	ret := make([][]int, a)
	for i := 0; i < a; i++ {
		ret[i] = make([]int, b)
	}
	return ret
}

// 整数
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}
func gcd(a, b int) int {
	c := 1
	for c != 0 {
		c = a % b
		a, b = b, c
	}
	return a
}
func divisorList(n int) []int {
	var l []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			l = append(l, i)
			if i != n/i {
				l = append(l, n/i)
			}
		}
	}
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
	return l
}
func factorization(n int) map[int]int {
	m := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n = n / i
		}
	}
	if n != 0 && n != 1 {
		m[n]++
	}
	return m
}
func isPrime(n int) bool {
	c := int(math.Sqrt(float64(n)))
	for i := 2; i <= c; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primeList(n int) []int {
	if n < 2 {
		return nil
	} else if n == 2 {
		return []int{2}
	}
	l := make([]int, n+1)
	primes := make([]int, 0, n)
	for i := 4; i <= n; i += 2 {
		l[i] = 1
	}
	for i := 3; i <= n; i += 2 {
		if l[i] == 1 {
			continue
		}
		for j := i * 2; j <= n; j += i {
			l[j] = 1
		}
	}
	primes = append(primes, 2)
	for i := 3; i <= n; i += 2 {
		if l[i] == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}
func combination(n, r int) int {
	r = min(n-r, r)
	d := make([]int, r)

	for i := 0; i < r; i++ {
		d[i] = n - i
	}
	for i := 2; i <= r; i++ {
		ti := i
		for j := 0; j < r; j++ {
			g := gcd(d[j], ti)
			if g != 1 {
				ti /= g
				d[j] /= g
			}
			if ti == 1 {
				break
			}
		}
	}
	ret := 1
	for i := 0; i < r; i++ {
		ret *= d[i]
	}
	return ret
}
func combinationMod(a, b, mod int) int {
	t1, t2 := 1, 1
	for i := 2; i <= b; i++ {
		t2 = (i * t2) % mod
	}
	gyaku := powMod(t2, mod-2, mod)

	for i := a - b + 1; i <= a; i++ {
		t1 = (i * t1) % mod
	}
	return ((t1 * gyaku) % mod)
}
func pow(a, b int) int {
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = (p * p)
			q = q / 2
		} else {
			t = (t * p)
			p = (p * p)
			q = (q - 1) / 2
		}
	}
	return (t * p)
}
func powMod(a, b, mod int) int {
	p, q := a, b
	t := 1
	for q >= 2 {
		if q%2 == 0 {
			p = (p * p) % mod
			q = q / 2
		} else {
			t = (t * p) % mod
			p = (p * p) % mod
			q = (q - 1) / 2
		}
	}
	return (t * p) % mod
}
func decimalToBaseN(s string, base int) string {
	n, _ := strconv.Atoi(s)
	t := strconv.FormatInt(int64(n), base)
	return t
}
func basenNToDecimal(s string, base int) string {
	t, _ := strconv.ParseInt(s, base, 64)
	ret := strconv.Itoa(int(t))
	return ret
}

// ビット操作
func mkBitList(x, keta int) []int {
	l := make([]int, keta)
	idx := 0
	for x != 0 {
		l[idx] = x & 1
		x = x >> 1
		idx++
	}
	return l
}

// Heap
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
