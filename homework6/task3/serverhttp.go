package task3

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func fillTemplate() {
	tmpl, err := template.ParseFiles("task3/template.html")
	if err != nil {
		log.Fatal(err)
	}

	data := TodoPageData{
		PageTitle: "TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(os.Stdout, data)

}

func tcpServerExample() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		defer conn.Close()
		for {
			_, err = io.WriteString(conn, "Hello tcp-world!\r\n")
			if err != nil {
				log.Print(err)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}
}

//Hello - get request from http server
//input in brouser example - http://localhost:8080/hello?name=Eva
func Hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	name := req.FormValue("name")
	log.Println(name)

	//fmt.Fprintf(res, "hello %v\n", name)
	io.WriteString(res,
			`<doctype html>
	<html>
	    <head>
	        <title>Hello World!</title>
	    </head>
	    <body>
	        Hello World*****!
	    </body>
	</html>`)
}

//Login - form 
//input login,pass in http://localhost:8080/login, return todo list for user
func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Println("method:", req.Method) //get request method
	if req.Method == "GET" {
		t, _ := template.ParseFiles("task3/login.gtpl")
		t.Execute(res, nil)

	} else {
		req.ParseForm()
		// logic part of log in
		name := req.Form["username"]
		pass := req.Form["password"]
		fmt.Println("username:", name)
		fmt.Println("password:", pass)

		//fmt.Fprintf(res, "hello %v\n", name)
		//-----------------------------------------------
		tmpl, err := template.ParseFiles("task3/template.html")
		if err != nil {
			log.Fatal(err)
		}
   
		todo:="todo list for user"+fmt.Sprint(name)
		
		data := TodoPageData{
			PageTitle: todo,
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(res, data)

	}
}
func HttpServer(){
	fs := http.FileServer(http.Dir("task3"))
	http.Handle("/", fs)             //http://localhost:8080
	http.HandleFunc("/hello", Hello) //http://localhost:8080/hello
	http.HandleFunc("/login", Login)
	
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("Bye")
	
}
//Task3 - start http server
func Task3() {
	HttpServer()
}
