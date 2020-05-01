package task1

import (
	"fmt"
)

//Cow - корова
type Cow struct{
	Animal
	Milk float32
}
//Speek - What does it say
func (animal *Cow) Speek(){
	fmt.Println(animal.Name," speek ",animal.Word)
}
//Eat - What does it eat
func (animal *Cow) Eat(){
	fmt.Println(animal.Name," eat ",animal.Food)
}
//Areal - What does it live
func (animal *Cow) Areal(){
	fmt.Println(animal.Name," areal ",animal.ArealStr)
	// return 2
}

