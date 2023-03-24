package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)

	// 魔法の回数 ＝ log(Nに含まれる素数の数)
	// Nに含まれる素数の数
	var count int64 = 0
	remain := n
	var i int64
	for i = 2; i*i <= n; i++ {
		for remain%i == 0 {
			count++
			remain /= i
		}
	}
	if remain != 1 {
		count++
	}

	// log(Nに含まれる素数の数)
	var ans int = 0
	var tmp int64 = 1
	for tmp < count {
		tmp *= 2
		ans++
	}

	fmt.Println(ans)
}
