package task4

import (
	"errors"
	"math"
)

var errD error=errors.New("ERR: D<0")
var errZero error=errors.New("ERR: /0")

func QuadraticEquation(a, b, c float64) (float64, float64, error) {
	d := b*b - 4*a*c
	var x1, x2 float64

	if d < 0 { //"Дискриминант меньше 0, корни невещественные."
		return x1, x2, errD
	}
	if(a==0){//Деление на 0
		return 0,0,errZero
	}
	//Если дискриминант больше или равен 0
	x1 = (-1*b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	x2 = (-1*b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	return x1, x2, nil
}
