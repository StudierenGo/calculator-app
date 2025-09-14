package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("__Калькулятор индекса массы тела__")

	height, weight := getUserInputs()
	imt := calculateImt(height, weight)

	if imt < 16 {
		fmt.Println("У вас сильный дефицит массы тела")
	} else if imt >= 16 && imt < 18.5 {
		fmt.Println("У вас дефицит массы тела")
	} else if imt >= 18.5 && imt < 25 {
		fmt.Println("У вас нормальный индекс массы тела")
	} else if imt >= 25 && imt < 30 {
		fmt.Println("У вас избыточная масса тела")
	} else if imt >= 30 && imt < 35 {
		fmt.Println("У вас 1-я степень ожирения")
	} else if imt >= 35 && imt < 40 {
		fmt.Println("У вас 2-я степень ожирения")
	} else {
		fmt.Println("У вас 3-я степень ожирения")
	}

	outputResult(imt)
}

func getUserInputs() (height float64, weight uint8) {
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&height)

	fmt.Println("Введите свой вес в килограммах: ")
	fmt.Scan(&weight)

	return
}

func calculateImt(height float64, weight uint8) float64 {
	const EXP = 2

	return float64(weight) / math.Pow(height/100, EXP)
}

func outputResult(imt float64) {
	fmt.Printf("Ваш индекс массы тела равен: %.0f", imt)
}
