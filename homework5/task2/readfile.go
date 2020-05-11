package task2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//ReadAll -  Load it entirely and read the file
func ReadAll()  {
	file, err := os.Open("task.txt")
    if err != nil {
        return
    }
    defer file.Close()

    // getting size of file
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // reading file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    fmt.Println(string(bs))

}

//ReadByChunks - Read file by chunks 
func ReadByChunks()  {
	file, err := os.Open("task.txt")
    if err != nil {
        fmt.Println("Error opening file!!!")
    }
    defer file.Close()
 
    // declare chunk size
    const maxSz = 3
 
    // create buffer
    b := make([]byte, maxSz)
 
    for {
        // read content to buffer
        readTotal, err := file.Read(b)
        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
        fmt.Print(string(b[:readTotal])) // print content from buffer
    }
}
// LinebyLineScan - read line by lines
// Note however bufio.Scanner has max buffer size 64*1024 bytes which means in case you file 
//has any line greater than the size of 64*1024, then it will give the error 
func LinebyLineScan() {
    file, err := os.Open("task.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
//Task2 - ответ на вопрос "Что исправить в программе чтения из файла"
func Task2()  {
	fmt.Println("HW 5 Task2")
	//1)Вариант чтения из лекции/презентации
	//Этот вариант имеет место только для маленьких файлов
	ReadAll() 

	//2)Варианты чтения больших файлов
	ReadByChunks()
	LinebyLineScan()

}