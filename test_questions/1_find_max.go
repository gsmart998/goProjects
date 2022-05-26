// 1.1 Найти максимальное число в массиве

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// generate random array of ints
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("test array:", arr)

	// find max value of array
	max := findMax(arr)

	fmt.Println("max value:", max)
}

// find max value using 'cycle i := 0'
func findMax(arr [10]int) (max int) {
	max = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
