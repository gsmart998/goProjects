package main

import "fmt"

func main() {
	fmt.Println("Hello, array!")

	planets := [...]string{ // Компилятор Go подсчитывает элементы
		"Меркурий",
		"Венера",
		"Земля",
		"Марс",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун", // Запятая в самом конце является обязательной
	}

	fmt.Println(planets)
	fmt.Println("Размер массива planets составляет:", len(planets))

	// Итерация через массивы
	dwarfs := [5]string{"Церера", "Плутон", "Хаумеа", "Макемаке", "Эрида"}
	// for i := 0; i < len(dwarfs); i++ {
	// 	dwarf := dwarfs[i]
	// 	fmt.Println(i, dwarf)
	// }
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

}
