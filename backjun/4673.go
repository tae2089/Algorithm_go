package backjun

import "fmt"

func d(n int) int {
	sum := n
	for true {
		if n == 0 {
			break
		}
		sum += n % 10
		n = n / 10
	}
	return sum
}

func function4673() {
	arr := make([]int, 100001)

	for i := 0; i < 10000; i++ {
		idx := d(i)
		if idx <= 10000 {
			arr[idx] = 1
		}
	}

	for i := 0; i < 10000; i++ {
		if arr[i] == 0 {
			fmt.Println(i)
		}
	}
}
