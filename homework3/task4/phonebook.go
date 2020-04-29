package task4
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)
//Person - структура человек с уникальным ID
type Person struct {
	ID      int
	Name    string
	Address string
}
//Account - связанная с человеком информация
type Account struct {
	Person
	IDAcc      int
	Owner   Person
	Business string
	Driverslicense int64
}
//Save - сохраняем значение структуры в файл
func (acc Account) Save(){
	b, err := json.Marshal(acc)
    if err != nil {
        log.Println(err)
        return
    }

	fmt.Println(b)
	fmt.Println(string(b))
	writeFile(File,b)
}

//File - имя файла, в которое сохраняем 
var File string="/tmp/phonebook.txt"
var id int=-1
func getID() int{
	id++
	return id
}

func writeFile(file string,message []byte)  {
	err := ioutil.WriteFile(file, message, 0644)
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERR :",err)
	}
}

//Task4 -задание
func Task4()  {
	fmt.Println("Task4")
	bb := Person{
		ID: getID(),
		Name:    "Dmitry",
		Address: "Moscow",
	}

	var acc Account = Account{
		IDAcc:     12,
		Owner:  bb,
		Driverslicense: 16658,
		Person : Person{getID(),"Mu","Main street"},
		Business: "Taxi",
	}
	fmt.Printf("%#v\n", acc)
	acc.Save()
} 

