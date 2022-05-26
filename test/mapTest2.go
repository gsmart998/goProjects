package main

import "fmt"

func main() {
	result := make([]int, 0)
	nums := []int{1, 1, 2, 8, 4, 3, 2, 3, 5, 5, 7, 7, 9}
	hashMap := make(map[int]int)

	// write to map numbers from array 'nums' with number of iterations
	for _, val := range nums {
		hashMap[val] = hashMap[val] + 1
	}
	fmt.Println(hashMap)

	// // write values from array that repeat once in map
	// for _, num := range nums {
	// 	if hashMap[num] == 1 {
	// 		result = append(result, num)
	// 	}
	// }

	// write values from array that repeat once in map
	for key := range hashMap {
		if hashMap[key] == 1 {
			result = append(result, key)
		}
	}

	// show array with unique values
	fmt.Println(result)
}
