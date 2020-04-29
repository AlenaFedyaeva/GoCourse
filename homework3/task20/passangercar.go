package task20
import (
	"encoding/json"
	"fmt"
	"log"
)

//PassangerCar - Пассажирская машина
//Цвет, количество сидений
type PassangerCar struct{
	CarInfo Car
	Color string
	NumberOfSeats int
}
//SetVal - задаем значения легковому автомобилю
func(mycar *PassangerCar)SetVal(car *Car,color string,numberOfSeats int){
	mycar.CarInfo.copyCar(*car)
	mycar.Color=color
	mycar.NumberOfSeats=numberOfSeats
}

//PrintPassInfo - выводим информацию на экран
func (mycar PassangerCar)PrintPassInfo()  {
	fmt.Println("PassangerCar")
	b, err := json.Marshal(mycar)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string(b))
}
