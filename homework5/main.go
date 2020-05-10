package main

import (
	"fmt"
	"time"
)


func main() {
	
	t := time.Now().UTC()
	t2 := time.Now()

	fmt.Println(t,"\n",t2)
	fmt.Println("\nThe end")
	
}
