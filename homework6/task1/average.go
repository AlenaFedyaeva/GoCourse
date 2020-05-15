package task1

import (
	"fmt"
)

func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total / float64(len(xs))
}

func SumArgs(arr ...int) (int) {
    sum:=0
    fmt.Println("len",len(arr))
    if(len(arr)==0){
        fmt.Println("Передано 0 элементов")
    }
    for _, a:= range arr {
        sum+=a
    }
    fmt.Println("sum=",sum)
    return sum
}