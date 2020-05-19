package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func mirroredQuery() string {
	var wg sync.WaitGroup
	var urls = []string{
        "http://github.com",
		"http://yandex.ru",
		"https://geekbrains.ru",
        // "http://google.com",
	}
	responses := make(chan string, 3)
	for _, url := range urls {
        // увеличиваем счетчик WaitGroup
		wg.Add(1)
		go func (url string)  {
			defer wg.Done()
			responses <- request(url)
		}(url)
		
    }
    // ожидаем, пока все запросы не будут завершены
    wg.Wait()
	
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) (string) {
	ServeHTTP(hostname)
	return hostname
}
func ServeHTTP(url string) bool {
	fmt.Println("serve ",url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	return true
}
func main() {
	fast:=mirroredQuery() 
	fmt.Println("fast: ",fast)

}
