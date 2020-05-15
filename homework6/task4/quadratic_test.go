package task4

import (
	"fmt"
	"testing"
)


type testVals struct {
    values []float64
	x1,x2 float64
	err error
}

var tests = []testVals {
    {[]float64{2,3,4}, 0,0,errD},
    {[]float64{1,-6,9},3,3,nil},
	{[]float64{0,1,1}, 0,0,errZero},
	{[]float64{0,1}, 0,0,errZero},
}


func TestQuadraticEquation(t *testing.T) {
    fmt.Println("QuadraticEquation")
	for _, val := range tests{
		if(len(val.values)<3){
			fmt.Println("ERR TEST: Test value length")
			continue
		}
		QuadraticEquation(val.values[0],val.values[1],val.values[2] )
	}
}
