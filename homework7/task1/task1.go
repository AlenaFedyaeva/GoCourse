package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
    for {
        for _, r := range "-\\|/" {
            fmt.Printf("%c\r", r)
            time.Sleep(delay)
        }
    }
}

func fibonacci(x int) int {
    if x < 2 {
        return x
    }
    return fibonacci(x-1) + fibonacci(x-2)
}

func main()  {
	go spinner(50 * time.Millisecond)
    const n = 42
	
	//Заставим спиннер вращаться
	//1) Способ 1
	time.Sleep(10*time.Second)

	fmt.Println("Bye")
}

