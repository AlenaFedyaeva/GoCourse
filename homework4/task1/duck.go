package task1

import (
	"fmt"
)

//Duck - утка
type Duck struct{
	Animal
	Fly bool
}

//Speek - What does it say
func (animal *Duck) Speek(){
	fmt.Println(animal.Name," speek ",animal.Word)
}
//Eat - What does it eat
func (animal *Duck) Eat(){
	fmt.Println(animal.Name," eat ",animal.Food)
}
//Areal - What does it live
func (animal *Duck) Areal(){
	fmt.Println(animal.Name," areal ",animal.ArealStr)
	// return 2
}

