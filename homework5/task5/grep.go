package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Extend(slice []int, element int) []int {
	if cap(slice) == len(slice) {
		fmt.Println("slice is full!", element)
		newSlice := make([]int, len(slice)+1) //, 2*cap(slice))
		copy(newSlice, slice)                 //копируем старые значения
		newSlice[len(slice)] = element        //добавляем новое
		slice = newSlice                      //заменяем слайс на новый
		return slice
	}
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

type FileType int

const (
	File = iota
	Dir
	NoneFile
)

func checkFile(name string) (bool, FileType) {
	file, err := os.Open(name)

	if err != nil {
		log.Println(err)
		return false,NoneFile
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Println(err)
		return false,NoneFile
	}
	if fi.IsDir() {
		return true,Dir
	} 
	return true,File
	
}

func printTypeFile(t FileType) (string) {
	switch t{
	case File:
		return "File"
	case Dir:
		return "Directory"
	default:
		return "NONE"
	}


}
//find - ищет первое вхождение в файл
func findInFile(word,fname string) bool{
	file, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		str:=scanner.Text()
		contains:=strings.Contains(str,word)
		if contains{
			return true
		}
        //fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return false
}

func grep(word,path string)  {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			_,typeF:=checkFile(path)
			if(typeF==File){
				bFind:=findInFile(word,path)
				if bFind {
					fmt.Println(path, info.Size(),printTypeFile(typeF),"Contain word ",word)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
//Сheck - reports whether the named file or directory exists.
func Сheck(fname string) bool {
	if _, err := os.Stat(fname); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
//Task5 - grep
func Task5() {
	//default values
	word:="main"
	path:="."

	flag.Parse()
	tail := flag.Args()
	fmt.Println("tail:", tail)


	switch len(tail) {
	default:
		fmt.Println("cp: too many arguments")
		return
	case 0:
		fmt.Println("cp: missing word & file operand")
		return
	case 1:
		fmt.Println("cp: missing file operand")
		return
	case 2:
		if !Сheck(tail[1]) {
			fmt.Println("cp: ", tail[1], " No such file/dir")
			return
		}
	}
	word=tail[0]
	path=tail[1]

	//Ищем файл, который содержит искомое слово
	grep(word,path)
}

func main() {
	Task5()
	fmt.Println("\nThe end")

}
