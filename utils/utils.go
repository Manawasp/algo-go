package utils

import "math/rand"

// RandArray generate an array of the N size with rand int
func RandArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
