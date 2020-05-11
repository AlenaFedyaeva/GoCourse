package main

import (
	"GoCource/homework5/task2"
	"fmt"
)


func Extend(slice []int, element int) []int {
	if cap(slice) == len(slice) {
		fmt.Println("slice is full!",element)
		newSlice := make([]int, len(slice)+1)//, 2*cap(slice))
		copy(newSlice, slice)//копируем старые значения
		newSlice[len(slice)]=element//добавляем новое
		slice = newSlice//заменяем слайс на новый
		return slice
	}
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	task2.Task2()
	

	fmt.Println("\nThe end")

}
