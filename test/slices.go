package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("new line")
	a := [...]int{
		8,
		2,
		64,
		3,
	}
	fmt.Println(a)
	var max = 0
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	fmt.Println("max value is:", max)

	var min = 0
	for i := 0; i < len(a); i++ {

		if a[i] < min {
			min = a[i]
		}

	}
	fmt.Println("min value is:", min)

	sort.Ints(a[:])
	fmt.Println(a)

}
