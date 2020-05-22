package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn, quit chan struct{}) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		fmt.Println("time", time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)

	}
}

func exit(quit chan struct{}) { //, wg *sync.WaitGroup) {
	for {
		str := ""
		fmt.Println("Для выхода наберите exit и нажмите кнопку Enter")
		fmt.Scanln(&str)
		if str == "exit" {
			fmt.Println("Bye!")
			quit <- struct{}{} 
			return
		}
	}
}
func tcpServer(quit chan struct{}) {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn, quit)
	}
}

func main() {
	quit := make(chan struct{})
	go exit(quit)
	go tcpServer(quit)
	<-quit
	close(quit)
}
