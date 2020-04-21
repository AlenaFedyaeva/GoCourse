package main

import (
	"math"
	"fmt"
	"os"
)

func banc(bankDeposit float64, percent float64){
	const years int =5
	fmt.Println("Программа для определения суммы вклада через 5 лет")
	fmt.Println("Введите сумму вклада рублях:")
	_,err:=fmt.Scanf("%f", &bankDeposit)
	if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
	}
	fmt.Println("Введите годовой процент:")
	_,err=fmt.Scanf("%f", &percent)
	if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
	}
	
	for i:=0;i<years;i++{
		yearIncome:=bankDeposit*percent/100
		bankDeposit=bankDeposit+yearIncome;
		fmt.Println("Год",i,"Проценты за год", math.Round(yearIncome),"Капитализация",math.Round(bankDeposit))
	}
	fmt.Println("Итого: ", bankDeposit)

}

func init() {
	fmt.Println("task3")
}
func main() {
	// Пользователь вводит сумму вклада в банк и годовой 
	// процент. Найти сумму вклада через 5 лет.
	banc(900,6);
	fmt.Println("The end")
}
