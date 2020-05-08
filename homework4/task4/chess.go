package task4
import (
	"fmt"
)

const boardLen int = 8

type Board struct {
	board [boardLen][boardLen]int
	posI int
	posJ int
}

var board Board
var possibleMoves [][]int
func init ()  {
	fmt.Println("Homework 4")
	board.setPos(0,0)
	possibleMoves=[][]int{{-1,-2},{-2, -1},{-2,  1},{ 1, -2},{-1,  2},{ 2,-1},{ 1, 2},{ 2, 1},}	
	//fmt.Print(possibleMoves)
}
func (b *Board) checkPos(i,j int) (bool){
	li:=len(b.board)
	lj:=len(b.board[0])
	if((i>=0 && i<li)&&(j>=0 && j<lj)){
		return true	
	}
	return false
}

func (b *Board) getMoves() ([][]int) {
	var arrMoves [][]int 
	for _, moves := range possibleMoves{
		i:=b.posI+moves[0]
		j:=b.posJ+moves[1]
		if(b.checkPos(i,j) && b.board[i][j]==0){
			arrMoves=append(arrMoves,[]int{i,j})
			//fmt.Println(i,j)
		}
	}
	fmt.Println("Can move to:",arrMoves)
	return arrMoves
}

func (b *Board) getPos() (int,int){
	fmt.Println("Position on board: ",b.posI,b.posJ)
	return b.posI,b.posJ
}
func (b *Board) printBoard() (){
	for _, arr := range b.board {
		fmt.Println(arr)
	}
}

func (b *Board) setPos(i,j int) (bool){
	li:=len(b.board)
	lj:=len(b.board[0])
	if(i<li && j<lj){
		b.board[i][i]=1
		b.posI=i
		b.posJ=j
		fmt.Println("Position on board changed to: ",b.posI,b.posJ)
	} else{
		return false
	}
	b.printBoard()
	return true
}


func Task4() {
	fmt.Println("Task 4")
	board.getPos()
	for i := 0; i < 5; i++ {
		arr:=board.getMoves()
		if(len(arr)>0){
			board.setPos(arr[0][0],arr[0][1])
		}
	}
	board.getMoves() 

	
}