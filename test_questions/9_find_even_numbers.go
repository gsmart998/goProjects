// 9. Найти в массиве четные числа

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// generate random array of ints
	var arr [10]int
	rand.Seed(time.Now().UnixNano()) //randomize numbers
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("test array:", arr)

	// find even numbers in array
	var result []int
	for _, val := range arr {
		if val%2 == 0 {
			result = append(result, val)
		}
	}

	// show even numbers
	fmt.Println("even numbers:", result)
}
