// 6. Найти среднее арифметическое из массива чисел

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

	// sum the numbers from array
	var sum int
	for _, val := range arr {
		sum = sum + val
	}

	// convert sum type to float64
	sum64 := float64(sum)

	// convert array len to float64
	arrayLen := float64(len(arr))

	// calculate avg number
	result := sum64 / arrayLen
	fmt.Println("avg numver:", result)
}
