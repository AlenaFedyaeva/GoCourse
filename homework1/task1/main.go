package main

import (
	"fmt"
)

func main() {
	//	Написать программу для конвертации рублей в доллары.
	// Программа запрашивает сумму в рублях и выводит сумму
	//в долларах. Курс доллара задайте константой.

	var rubles float64
	const course float64 = 70.04

	fmt.Println("Программа для конвертации рублей в доллары")
	fmt.Println("Введите сумму в рублях:")
	fmt.Scanf("%f", &rubles)
	fmt.Println("Вы ввели: ",rubles)
	fmt.Println("Расчитываем по курсу $:",course)
	dollars := rubles/course;
	fmt.Println("Итого долларов:", dollars)
	fmt.Println("The end")

}
