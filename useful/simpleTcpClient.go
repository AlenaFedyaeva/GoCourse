package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    io.Copy(os.Stdout, conn) // игнорируем ошибки
}
