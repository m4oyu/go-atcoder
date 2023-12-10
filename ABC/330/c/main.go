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
	const buf = 200100
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, buf), buf)
}

func debug() {
	input, err := os.Open("./in")
	if err != nil {
		fmt.Fprintln(os.Stderr, "there is no input file")
		os.Exit(1)
	}
	sc = bufio.NewScanner(input)
}

func main() {
	sc.Scan()
	d, _ := strconv.ParseFloat(sc.Text(), 64)

	r := int(math.Ceil(math.Sqrt(d)))

	ans := 2 * math.Pow10(12)

	for i := 0; i < r; i++ {
		a := d - math.Pow(float64(i), 2)
		y := math.Sqrt(a)
		yCeil := math.Abs(a - math.Pow(math.Ceil(y), 2))
		if yCeil < ans {
			ans = yCeil
		}
		yFloor := math.Abs(a - math.Pow(math.Floor(y), 2))
		if yFloor < ans {
			ans = yFloor
		}
	}

	fmt.Println(int(ans))

}
