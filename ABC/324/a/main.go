package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a := strings.Split(sc.Text(), " ")

	for i := 0; i < n-1; i++ {
		if a[i] != a[i+1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}
