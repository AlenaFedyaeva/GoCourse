package task20
import (
	"encoding/json"
	"fmt"
	"log"
)
//Truck - Машина типа грузовик
//Тоннаж, тип грузов
type Truck struct{
	CarInfo Car
	Cargo  TypeOfCargo  `json:"cargo"`
}

//TypeOfCargo - тип груза 
type TypeOfCargo int
const (
    Oil TypeOfCargo = iota
    Milk TypeOfCargo = iota
    Sand TypeOfCargo = iota
)

func (a TypeOfCargo) MarshalJSON() ([]byte, error) {
	var s string
	switch a{
		case Oil: s="oil"
		case Milk: s="milk"
		case Sand: s="sand"
		default: s="NONE"
		}

	return json.Marshal(s)
}

// SetVal - задаем значения грузовику
func(truck *Truck)SetVal(car Car,cargo TypeOfCargo){
	truck.CarInfo.copyCar(car)
	truck.Cargo=cargo
}

//PrintTruckInfo - выводим информацию на экран
func (truck Truck)PrintTruckInfo()  {
	fmt.Println("Truck")
	b, err := json.Marshal(truck)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string(b))
}



