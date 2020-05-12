package main

import (
	// "GoCource/homework5/task2"
	// "GoCource/homework5/task3"
	"flag"
	"fmt"
	"os"
)

func Extend(slice []int, element int) []int {
	if cap(slice) == len(slice) {
		fmt.Println("slice is full!", element)
		newSlice := make([]int, len(slice)+1) //, 2*cap(slice))
		copy(newSlice, slice)                 //копируем старые значения
		newSlice[len(slice)] = element        //добавляем новое
		slice = newSlice                      //заменяем слайс на новый
		return slice
	}
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
//CpUtil - утилита копирования файла
func CpUtil()  {
	argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[0]

	if(len(argsWithoutProg)==0){
		fmt.Println("cp: missing file operand\nTry 'cp --help' for more information.")
		return
	}

    fmt.Println("withpro",argsWithProg)
    fmt.Println("argsWithout",argsWithoutProg)
    fmt.Println("arg",arg)

	

	srcPtr := flag.String("src", "", "file src name")
    dstPtr := flag.String("dst", "", "file dst name")

    flag.Parse()

    fmt.Println("src:", *srcPtr)
    fmt.Println("dst:", *dstPtr)
  
    fmt.Println("tail:", flag.Args())


}

func main() {
	// task2.Task2()
	// task3.Task3()

	CpUtil()


	fmt.Println("\nThe end")

}
