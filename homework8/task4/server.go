
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

type MessageType int
const(
	Name MessageType=iota
	Auth
)

type ConnInfo struct{
	name string
	port string
	ip string
}

var(
	mu sync.Mutex
	connArr []ConnInfo
)

func getConnInfo(ip,port string) (name string) {
	mu.Lock()
	for _, conn := range connArr {
		if(conn.port==port && conn.ip==ip){
			name=conn.name
			break
		}
	}
	mu.Unlock()
	return name
}
func addConnInfo(ip,port,name string){
	mu.Lock()
	connArr=append(connArr,ConnInfo{name,port,ip})
	mu.Unlock()
	
}


type client chan<- string

var (
    entering = make(chan client)
    leaving  = make(chan client)
    messages = make(chan string)
)

func main() {
	
    listener, err := net.Listen("tcp", "localhost:8090")
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Server start",listener.Addr().Network(),listener.Addr().String())
    go broadcaster()
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
		}
		
        go handleConn(conn)
    }
}
func broadcaster() {
    clients := make(map[client]bool)
    for {
        select {
		case msg := <-messages:
			fmt.Println(">>> ",msg)
            for cli := range clients {
                cli <- msg
            }

        case cli := <-entering:
            clients[cli] = true

        case cli := <-leaving:
            delete(clients, cli)
            close(cli)
        }
    }
}

func handleConn(conn net.Conn) {
    ch := make(chan string)
    go clientWriter(conn, ch)

	name:=""

	input := bufio.NewScanner(conn)
	b:=true
    for input.Scan() {
		msg:=input.Text()
		if(b){
			b=false
			args:=strings.Split(msg, "|") 
			name=args[1]
			addConnInfo(conn.RemoteAddr().String(),"", name)
			ch <- "You are " + name
			messages <- name + " has arrived"
			entering <- ch
		}else{
			 messages <- name + ": " + msg
		}
       
    }
    leaving <- ch
    messages <- name + " has left"
    conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
    for msg := range ch {
        fmt.Fprintln(conn, msg)
    }
}
