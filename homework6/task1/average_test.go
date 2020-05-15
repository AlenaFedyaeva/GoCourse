package task1

import (
	"fmt"
	"log"
	"testing"
)

type testpair struct {
    values []float64
    average float64
}

var tests = []testpair {
    {[]float64{1,2}, 1.5},
    {[]float64{1,1,1,1,1,1}, 1},
    {[]float64{-1,1}, 0},
}

func TestAverage(t *testing.T) {
    fmt.Println("TestAverage!!!")
    for _, pair := range tests {
        v := Average(pair.values)
        if v != pair.average {
            t.Error(
                "For", pair.values, 
                "expected", pair.average,
                "got", v,
            )
        }
    }
}


type testSum struct {
    values []int
    sum int
}

var testSumArr = []testSum {
    {[]int{1,2}, 3},
    {[]int{1,1,1,1,1,1}, 6},
    {[]int{-1,1}, 0},
    {[]int{}, 0},
}



func TestSumArgs(t *testing.T) {
    fmt.Println("TestSumArgs!!!")
    for _, pair := range testSumArr{
        log.Println(pair,pair.values)
        v := SumArgs(pair.values...)
        if v != pair.sum {
            t.Error(
                "For", pair.values, 
                "expected", pair.sum,
                "got", v,
            )
        }
    }
}