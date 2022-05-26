// 2. Из массива вычленить только уникальные значения

package main

import "fmt"

func main() {
	result := make([]int, 0)
	nums := []int{1, 1, 2, 8, 4, 3, 2, 3, 5, 5, 7, 7, 9}
	hashMap := make(map[int]int)

	fmt.Println("test array:", nums)

	// write to map numbers from array 'nums' with number of iterations
	for _, val := range nums {
		hashMap[val] = hashMap[val] + 1
	}

	// write values from array that repeat once in map
	for key := range hashMap {
		if hashMap[key] == 1 {
			result = append(result, key)
		}
	}

	// show array with unique values
	fmt.Println("unique values:", result)
}
