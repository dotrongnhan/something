package fibonacci

import "fmt"

func fib(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n - 1; i++ {
		var f int
		if i <= 2 && i > 0 {
			f = 1
		} else if i == 0 {
			f = 0
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fmt.Println(f)
		fn[i] = f
	}
	end, _ := fmt.Printf("kết thúc chuỗi với độ dài %x", n)
	return end
}
