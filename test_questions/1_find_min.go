// 1. Найти минимальное число в массиве

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

	// find min value of array

	min := findMin(arr)

	fmt.Printf("1. min value: %d \n", min)

}

// find min value using 'range'

func findMin(arr [10]int) (min int) {
	min = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
