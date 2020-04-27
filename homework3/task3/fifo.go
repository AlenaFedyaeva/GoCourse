package task3
import (
	"fmt"
	"strconv"
)


//Task3 - FIFO
func Task3()  {
	fmt.Println("Task 3 - FIFO")

	//Заполним значениями
	for i:= 0; i < 15; i++ {
		strIn:="next"+strconv.Itoa(i)
		strOut:="NONE"
		b,strOutFifo:=Fifo(strIn)
		if b {
			strOut=strOutFifo
			
		}
		fmt.Println("in - ",strIn, Queue," out ",strOut)
	}
}

var lenFifo int =3

//Queue - массив, хранящий FIFO
var Queue []string

//Fifo - фукция FIFO, возвращает true если есть выходное значение из буфера
func Fifo (inStr string)(isOutput bool,outStr string )  {
	if(len(Queue)==lenFifo){
		outStr=Queue[lenFifo-1]
		isOutput=true
		Queue = Queue[1:]   // Dequeue
		Queue = append(Queue, inStr)
	}else{
		Queue = append(Queue, inStr)
	}
	fmt.Println(Queue)
	return isOutput,outStr
}
