package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func handleConn(c net.Conn, quit chan struct{}) {
	defer c.Close()
	for {
		select {
		case <-quit:
			fmt.Println("hBYE")
			quit <-struct{}{} 
			return
		default:
			fmt.Println("hdefault")
			_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
			fmt.Println("time", time.Now().Format("15:04:05\n\r"))
			if err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}

	}
}

func exit(quit chan struct{}) { //, wg *sync.WaitGroup) {
	// defer wg.Done()
	for {
		str := ""
		fmt.Println("Для выхода наберите exit и нажмите кнопку Enter")
		fmt.Scanln(&str)
		if str == "exit" {
			fmt.Println("Bye!")
			quit <- struct{}{} // send sum to c
			return
		}
	}
}
func tcpServer(quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case <-quit:
			fmt.Println("BYE all")
			close(quit)
			return
		default:
			fmt.Println("default")
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			handleConn(conn, quit)
		}

	}
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	quit := make(chan struct{})
	defer close(quit)
	go exit(quit)
	go tcpServer(quit, wg)

	wg.Wait()
}
