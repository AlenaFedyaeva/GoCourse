package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const trackLen float32 = 10000.5

type auto struct {
	name      string
	speed     int
	trackTime float32
}

func (a *auto) autoInit(name string) {
	a.name = name
	a.speed = rand.Intn(350) //gen random speed
	a.trackTime = trackLen / float32(a.speed)
}

//ridesTrack - машина проезжает трек за время t=S/v
func (a *auto) ridesTrack() {
	time.Sleep(time.Duration(a.trackTime) * time.Millisecond)
}

func (a *auto) race(raceRez chan string,wg *sync.WaitGroup) {
	defer wg.Done()
	tStart := time.Now()
	a.ridesTrack()
	tEnd := time.Now().Sub(tStart)
	raceRez <- a.name
	fmt.Println("run ", a.name, "time track ", a.trackTime,"time rez ", tEnd)
}
var names = []string{"ford", "bmv", "subaru","mazda"}
var autoArr []auto 

func init(){
	fmt.Println("init cars")
	for _, name := range names {
		var a auto
		a.autoInit(name)
		autoArr=append(autoArr,a)
	}
	fmt.Println(autoArr)
}

func main() {
	wg:=&sync.WaitGroup{}
	raceRez:=make(chan string)

	for _, a := range autoArr {
		fmt.Println("go ",a)
		wg.Add(1)
		go a.race(raceRez,wg)
	}
	
	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		for name := range raceRez {
			fmt.Println(" place ",name)
		}
	}(wg)
	wg.Wait()
	close(raceRez)
	
}
