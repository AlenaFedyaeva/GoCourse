package task2

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

//NumberList - список номеров телефонов
type NumberList []int
func (n NumberList) Len() int           { return len(n) }
func (n NumberList) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n NumberList) Less(i, j int) bool { return n[i]< n[j]}


// Phones - тип, описывающий контакт в телефонной книге
type Phones struct{
	Name string
	Numbers NumberList//[]int
}
// PhoneBook - телефонная книга
type PhoneBook []Phones

//PrintAll - печать на экран
func (phoneBook PhoneBook) PrintAll()  {
	for i, phone := range phoneBook {
		fmt.Printf("\t %v) %#v %v\n",i,phone.Name,phone.Numbers )
	}
}
//Len - возвращает количество элементов в коллекции.
func (phoneBook PhoneBook) Len() int  {
	return len(phoneBook)
}
//Less - отдает true, если элементы i, j необходимо поменять местами.
func (phoneBook PhoneBook) Less(i,j int) bool  {
	iRune, _ := utf8.DecodeRuneInString(phoneBook[i].Name)
	jRune, _ := utf8.DecodeRuneInString(phoneBook[j].Name)
	return int32(iRune) < int32(jRune)
}
// Swap меняет местами элементы с индексами i,j.
func (phoneBook PhoneBook) Swap (i,j int){
	phoneBook[i],phoneBook[j]=phoneBook[j],phoneBook[i]
}

//Sort - функция для сортировки телефонной книги
//Записи сортируются по имени в алфавитном  
//Номера телефона в записи сортируются по возрастанию
func (phoneBook *PhoneBook) Sort(){
	sort.Sort(phoneBook)
	for _, book := range *phoneBook {
		sort.Sort(book.Numbers)
		//Можно и просто отсортировать встроенной функцией для Int
		//sort.Ints(book.Numbers)
	}
}

//Task2 - задание 2
func Task2()  {
	fmt.Println("Task 2 - adress book")
	phoneBook:=PhoneBook{Phones{"Max",[]int{53,89167253764,2,89635437382}},}
	phone:=Phones{"Alex",[]int{78293467382,456,}}
	phoneBook=append(phoneBook,phone)
	phoneBook=append(phoneBook,Phones{"Lena",[]int{654,78253464382,}})
	
	fmt.Println("1) Before sort()")
	phoneBook.PrintAll()
	fmt.Println("2) After sort() by Name")
	phoneBook.Sort()
	phoneBook.PrintAll()
}
