package main

import (
	"fmt"
	"math"
)

func hypotenuse(a float64,b float64) float64{
	res:=math.Sqrt(a*a+b*b);
	return res 
}

func perimeter(a float64,b float64) float64{
	c:=hypotenuse(a,b)
	res:=a+b+c;
	return res
}

func area(a float64,b float64) float64{
	res:=(a*b)/2;
	return res
}

func init() {
	fmt.Println("task2")
}
func main() {
	var cathetA float64 = 3
	var cathetB float64 = 4

	fmt.Println("Найдем площадь, периметр и гипотенузу треугольника по катетам")
	fmt.Println("Значения катетов ", cathetA, cathetB)
	
	fmt.Println("Hypotenuse",hypotenuse(cathetA,cathetB))
	fmt.Println("Perimeter",perimeter(cathetA,cathetB))
	fmt.Println("Area",area(cathetA,cathetB))
	fmt.Println("The end")
}
