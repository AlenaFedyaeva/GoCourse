package task1

import (
	"fmt"
)

//Dolphin - дельфин
type Dolphin struct{
	Animal
	Swim bool
}
//Speek - What does it say
func (animal *Dolphin) Speek(){
	fmt.Println(animal.Name," speek ",animal.Word)
}
//Eat - What does it eat
func (animal *Dolphin) Eat(){
	fmt.Println(animal.Name," eat ",animal.Food)
}
//Areal - What does it live
func (animal *Dolphin) Areal(){
	fmt.Println(animal.Name," areal ",animal.ArealStr)
	// return 2
}
