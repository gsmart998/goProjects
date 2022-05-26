package main

import (
	"errors"
	"fmt"
)

var winePrice = 100

func main() {
	change, err := buyWine(20, 130)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Ваша сдача - %d. Хорошего дня!\n", change)
	}
	change, err = buyWine(18, 90)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Ваша сдача - %d. Хорошего дня!\n", change)
	}
	change, err = buyWine(15, 130)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Ваша сдача - %d. Хорошего дня!\n", change)
	}
}

func buyWine(age, moneyAmount int) (int, error) {
	if age >= 18 && moneyAmount >= winePrice {
		return moneyAmount - winePrice, nil
	} else if age < 18 {
		return moneyAmount, errors.New("Иди пей вишневый сок!")
	} else {
		return moneyAmount, errors.New("У вас недостаточно средств!")
	}
}
