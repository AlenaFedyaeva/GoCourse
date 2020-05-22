package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	ch := make(chan string)
// 	args := os.Args[1:]
// 	if len(args) == 0 {
// 		log.Println("Задайте имя участника чата")
// 		return
// 	} else if len(args) > 1 {
// 		log.Println("Имя может состоять только из 1 слова")
// 		return
// 	}
// 	name := args[0]
// 	fmt.Println("My name in room:", name)

// 	conn, err := net.Dial("tcp", "localhost:8090")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	str := "name|" + name
// 	fmt.Fprintln(conn, str)

// 	// writer := bufio.NewWriter(conn)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	os.Exit(1)
// 	// }

// 	// writer.WriteString(str)  // запись строки
// 	// writer.WriteString("\n") // перевод строки
// 	// writer.Flush()           // сбрасываем данные из буфера в файл

// 	go func() {
// 		input := bufio.NewScanner(conn)
// 		for input.Scan() {
// 			msg := name + "|msg|" + input.Text()
// 			log.Println(msg)
// 		}
// 	}()

// 	go func() {
// 		input := bufio.NewScanner(os.Stdin)
// 		for input.Scan() {
// 			msg := name + ":" + input.Text()
// 			fmt.Fprintln(conn, msg)
// 		}
// 	}()
// 	<-ch
// 	close(ch)
// 	//io.Copy(conn, os.Stdin) // until you send ^Z
// 	fmt.Printf("%s: exit", conn.LocalAddr())
// }
