package task20

import (
	"encoding/json"
	"fmt"
	"log"
)

//Car - структура, содержащая общую информацию для автотранспорта
type Car struct{
	CarModel string
    Year int
    IsWindowOpen bool
    IsEnginOn bool
    CapacityOfTrunk int
    ZVolumeTrunk int
}

//SetVal - устанавливает значения в поля структуры Car
func (mycar *Car)SetVal(carModel string,year int,isWindowOpen bool,isEnginOn bool,capacityOfTrunk int,zVolumeTrunk int){
	mycar.CarModel=carModel
	mycar.Year=year	
	mycar.IsWindowOpen=isWindowOpen
	mycar.IsEnginOn=isEnginOn
	mycar.CapacityOfTrunk=capacityOfTrunk
	mycar.ZVolumeTrunk=zVolumeTrunk
}
//copyCar - устанавливает значения в поля структуры Car
func (mycar *Car)copyCar(car Car){
	mycar.CarModel=car.CarModel
	mycar.Year=car.Year	
	mycar.IsWindowOpen=car.IsWindowOpen
	mycar.IsEnginOn=car.IsEnginOn
	mycar.CapacityOfTrunk=car.CapacityOfTrunk
	mycar.ZVolumeTrunk=car.ZVolumeTrunk
}

//PrintCarInfo - выводим информацию на экран
func (mycar Car)PrintCarInfo()  {
	fmt.Println("Car")
	b, err := json.Marshal(mycar)
    if err != nil {
        log.Println(err)
        return
    }
    
    fmt.Println(string(b))
}