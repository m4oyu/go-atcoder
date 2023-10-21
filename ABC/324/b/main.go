package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for {
		if n%2 == 0 {
			n /= 2
			continue
		}
		if n%3 == 0 {
			n /= 3
			continue
		}
		break
	}

	if n == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
