package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ch := make(chan string)
	args := os.Args[1:]
	if len(args) == 0 {
		log.Println("Задайте имя участника чата")
		return
	} else if len(args) > 1 {
		log.Println("Имя может состоять только из 1 слова")
		return
	}
	name := args[0]
	fmt.Println("My name in room:", name)

	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	str := "name|" + name
	fmt.Fprintln(conn, str)

	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			msg := input.Text()
			log.Println(msg)
		}
	}()

	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			msg := input.Text()//here add some info into conn
			fmt.Fprintln(conn, msg)
		}
	}()
	<-ch
	close(ch)
	//io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}
