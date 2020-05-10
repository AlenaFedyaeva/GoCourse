package task4
import (
	"fmt"
	"math/rand"
)

const boardLen int = 8
//board -  доска
var board Board
//posibloPints -  - массив, содержащий значения точек для вычисления возможных движений коня на доске
var possiblePoints []Point

//Point - точка на шахматной доске
type Point struct{
	x int
	y int
}
func (p *Point) setValue(in Point)  {
	p.x=in.x
	p.y=in.y
}
//Board - шахматная доска 8*8
type Board struct {
	board [boardLen][boardLen]int
	Point
}


//checkPos - функция, проверяющая находится ли позиция на доске
func (b *Board) checkPos(i,j int) (bool){
	lj:=len(b.board)
	li:=len(b.board[0])
	if((i>=0 && i<li)&&(j>=0 && j<lj)){
		return true	
	}
	return false
}

//getPoints - возвращает массив из точек, в которые конь сможет сделать ход
//Проверяем находтся ли позиция на доске 
func (b *Board) getPoints() ([]Point) {
	var arr []Point 
	for _, pt := range possiblePoints{
		x:=b.x+pt.x
		y:=b.y+pt.y
		if(b.checkPos(x,y)){
			arr=append(arr,Point{x,y})
		}
	}
	fmt.Println("Can move to points:",arr)
	return arr
}

func (b *Board) getPos() (Point){
	fmt.Println("Position on board: ",b.x,b.y)
	return Point{b.x,b.y}
}

func (b *Board) printBoard() (){
	for i := 0; i < len(b.board[0])*2+3; i++ {
		if(i==0){
			fmt.Print("   ")
		}
		fmt.Print("-")
	}
	fmt.Print("x\n")
	for j, arr := range b.board {
		fmt.Println(j,"|",arr)
	}
	fmt.Println("  y ")
}
//setPoint - выставляет позицию коня на доску, в декартовых координатах (x,y). 
// Х - по горизонтали, У- по вертикали
//Начало координат (0,0) в левом верхнем углу доски
func (b *Board) setPoint(p Point) (bool){
	lj:=len(b.board)
	li:=len(b.board[0])
	if(p.x<li && p.y<lj){
		b.board[p.y][p.x]=1
		b.Point.setValue(p)
		fmt.Println("Position on board changed to point: ",b.Point)
	} else{
		return false
	}
	return true
}


func init ()  {
	fmt.Println("Task 4")
	board.setPoint(Point{0,0})
	possiblePoints=[]Point{{-1,-2},{-2, -1},{-2,  1},{ 1, -2},{-1,  2},{ 2,-1},{ 1, 2},{ 2, 1},}
}

//Task4 - задание 4
func Task4() {
	arr:=board.getPoints()
	board.setPoint(arr[0])
	for i := 0; i < 5; i++ {
		arr:=board.getPoints()
		if(len(arr)>0){
			rnd:=rand.Intn(len(arr))
			board.setPoint(arr[rnd])
			board.printBoard()
		}
	}
}