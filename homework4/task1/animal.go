package task1

import (
	"fmt"
)

//AnimalInterface - интерфейс Животные
type AnimalInterface interface{
	Speek()
	Eat() 
	Areal() 
}
//Animal - структура Животные
type Animal struct{
	Name string
	Age int
	Weight float32
	Food string
	Word string
	ArealStr string

}

//PrintAnimal - выводим информацию о животных
func PrintAnimal(animals ...AnimalInterface){
	for _, a := range animals {
		a.Speek()
		a.Areal()
		a.Eat()
	}
}
//Task1 - первое задание
func Task1()  {
	fmt.Println("Task1")
	cow:=&Cow{Animal{"Cow",2,100,"grass","Mu-u-u","Farm"},12}
	dolphin:=&Dolphin{Animal{"Dolphin",11,56,"fish","Pi-i-iii","Ocean"},true}
	duck:=&Duck{Animal{"Duck",1,5,"corn","Quack-quack","Forest"},true}
	PrintAnimal(cow,dolphin,duck) 

}