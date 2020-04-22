package packagehw2
import (
	"fmt"
	"math/big"
	"os"
)
//Task1 - проверяем введеное число на четность
func Task1() {
	var number int
	fmt.Println("Введите целое число для проверки на четность")
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Вы ввели: ",number)
	fmt.Println("Проверяем число ... ", EvenNumber(number))
}
//Task2 - проверяем делится ли введеное число без остатка на 3. 
func Task2()  {
	var number int
	fmt.Println("Введите целое число для проверки деления без остатка на 3")
	_, err := fmt.Scanf("%d\n", &number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println("Вы ввели: ",number)
	fmt.Println("Проверяем число ... ", DevisionOn3(number))
}
//Task3 - выводим на экран 100 первых чисел Фибоначчи, начиная от 0. 
func Task3(){
	const count int=100
	fibNumber :=big.NewInt(1)
	fibNumberPrevious :=big.NewInt(0)
	for i := 0; i < count; i++ {
		fibNumber.Add(fibNumber,fibNumberPrevious)
		fibNumber,fibNumberPrevious=fibNumberPrevious,fibNumber
		fmt.Print(" ,",fibNumber)
	}
}


//DevisionOn3 -  функция, которая определяет, 
//делится ли число без остатка на 3
func DevisionOn3(number int) string{
	if(number%3==0){
		return "делится на 3 без остатка"
	}
	return "не делится на 3 без остатка"
}

//EvenNumber - функция, которая определяет, четное ли число.
func EvenNumber(number int) string{
	if(number%2==0){
		return "четное"
	}
	return "нечетное"
}