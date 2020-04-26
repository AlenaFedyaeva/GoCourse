package task20
import (
	"fmt"
)
//Task12 - Задание 1 & 2
func Task12()  {
	fmt.Println("Задание 1 & 2")
	//1) Задаем значения сразу
	car1:=Car{"Volvo",1998,true,false,150,50}
	car1.PrintCarInfo()
	//2) Создаем со значениями по умолчанию
	car2:=Car{}
	car2.PrintCarInfo()
	//3) Используем метод для заполнения значений
	car2.SetVal("ford",2020,false,true,2000,10)
	car2.PrintCarInfo()
	//4) Машина типа грузовик
	truck:=Truck{}
	truck.SetVal(car1,Oil)
	truck.PrintTruckInfo()
	//5) Легковой автомобиль
	passangerCar:=PassangerCar{}
	passangerCar.SetVal(&car2,"red",12)
	passangerCar.PrintPassInfo()
}
