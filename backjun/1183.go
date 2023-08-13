package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	l := make([]int, n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		l[i] = a - b
	}

	sort.Ints(l)

	if n%2 == 1 {
		fmt.Println(1)
	} else {
		fmt.Println(l[n/2] - l[n/2-1] + 1)
	}
}
