package main

import (
	"fmt"
)

//конвеер
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	const num int = 50
	// генерация
	go func() {
		for x := 0; x < num; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for val := range squares {
		fmt.Println(val)
	}
}
