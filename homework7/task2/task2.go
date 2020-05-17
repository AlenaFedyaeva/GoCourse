package main

import (
	"fmt"
)

//конвеер
func main() {
    naturals := make(chan int)
    squares := make(chan int)

	const num int=50
    // генерация
    go func() {
        for x := 0;x<num; x++ {
            naturals <- x
		}
		close(naturals)
    }()

    // возведение в квадрат
    go func() {
        for {
			x, ok := <-naturals
			if(!ok){
				close(squares)
				return
			}
            squares <- x * x
			
        }
    }()

    // печать
    for {
		_, ok := <-squares
		if !ok {
			break
		}
        fmt.Println(<-squares)
    }
}