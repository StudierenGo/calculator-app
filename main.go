package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")

	for {
		height, weight := getUserInputs()
		imt, error := calculateImt(height, weight)

		if error != nil {
			fmt.Println("Не задан рост или вес")
			continue
		}

		outputResult(imt)

		if !checkUserAction() {
			break
		}
	}
}

func getUserInputs() (height float64, weight uint8) {
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)

	fmt.Println("Введите свой вес в килограммах: ")
	fmt.Scan(&weight)

	return
}

func calculateImt(height float64, weight uint8) (float64, error) {
	const EXP = 2

	if height <= 0 {
		return 0, errors.New("NOT_PROVIDED_HEIGHT")
	}

	if weight <= 0 {
		return 0, errors.New("NOT_PROVIDED_WEIGHT")
	}

	return float64(weight) / math.Pow(height/100, EXP), nil
}

func outputResult(imt float64) {
	message := fmt.Sprintf("Ваш индекс массы тела равен: %.0f", imt)
	fmt.Println(message)

	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный индекс массы тела")
	case imt < 30:
		fmt.Println("У вас избыточная масса тела")
	case imt < 35:
		fmt.Println("У вас 1-я степень ожирения")
	case imt < 40:
		fmt.Println("У вас 2-я степень ожирения")
	default:
		fmt.Println("У вас 3-я степень ожирения")
	}
}

func checkUserAction() bool {
	userAction := "y"
	fmt.Print("Хотите продолжить (y/n)? ")
	fmt.Scan(&userAction)

	return strings.ToLower(userAction) == "y" || strings.ToLower(userAction) == "yes"
}
