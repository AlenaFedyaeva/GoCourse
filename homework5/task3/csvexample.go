package task3

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)
var csvFileName string="tmp.csv"

//WriteCsv - how to write into csv
func WriteCsv() {
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	//1) Запишем одну строку
	csvWriter := csv.NewWriter(csvFile)
	err= csvWriter.Write([]string{"678", "Jane", "jane@example.com", "$548,980"})
	//2) Запишем несколько строк
	rows := [][]string{
		{ "123", "John", "john@example.com", "$141,987"},
		{ "456", "Sam", "sam@example.com", "$905,234"},
		{ "678", "Jane", "jane@example.com", "$548,980"},
	  }
	  
	err = csvWriter.WriteAll(rows)
	//3) Сохраняем записанное в файл
	csvWriter.Flush()
	err= csvWriter.Error()
	if err != nil {
  	fmt.Println("an error occurred during the flush")
	}
}


//ReadCsv - example how to read from csv file
func ReadCsv() {
	csvFile, err := os.Open("tmp.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	
	csvReader := csv.NewReader(csvFile)

	//1) Читаем одну строку
	row, err := csvReader.Read()
	fmt.Println(row, row[1])
	//2) Читаем остальные строки
	for {
		row, err = csvReader.Read()
		if err == io.EOF {
		  break
		} else if err != nil {
		  panic(err) // or handle it another way
		}
		fmt.Println(row)
	  }
}


//Task3 - csv how to
func Task3()  {
	// 1) Запишем пример csv 
	WriteCsv()
	// 2) Прочитаем из файла
	ReadCsv()
}