// Тестовые вопросы:
// 1. Найти минимальное/максимальное число в массиве [готово]
// 2. Из массива вычленить только уникальные значения [готово]
// 3. Из массива вычленить только не уникальные значения [готово]
// 4. Обойти дерево (просто и бинарное)
// 5. Обойти связанный список
// 6. Найти среднее арифметическое из массива чисел
// 7. Решить задачу на валидацию последовательности скобок https://leetcode.com/problems/valid-parentheses
// 8. Решить задачу на палиндром https://leetcode.com/problems/valid-palindrome/
// 9. Найти в массиве четные числа

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Start")

	// generate random array of ints
	var arr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("test array is: ", arr)

	var result []int
	for _, val := range arr {
		if val%2 == 0 {
			result = append(result, val)
		}
	}
	fmt.Println(result)

	// 1. find min value of array

	min := findMin(arr)

	fmt.Printf("1. min value is: %d \n", min)

	// 2. find max value of array

	max := findMax(arr)

	fmt.Println("2. max value is: ", max)

	// 3. find unique value of array

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
