package checkPrime

import "math"

func CheckNumberPrime(num int) bool {

	if num < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPrime(arr []int) []int {
	var prime []int
	for i := 0; i < len(arr); i++ {
		if CheckNumberPrime(arr[i]) {
			prime = append(prime, arr[i])
		}
	}

	return prime
}
