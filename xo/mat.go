package main

import (
	"fmt"
)

func input[Loc any](prompt string , location  *Loc) {
	fmt.Print(prompt+": ")
	fmt.Scan(location)
}

type Matrix  struct{
	row,column uint 
	value  [][]int
}

func main() {
	var myMatrix Matrix
	myMatrix.getDimensions()
	myMatrix.value=func(row, col int) [][]int {
		return make([][]int,row)
	}
	myMatrix.getValues()

}

func (mat *Matrix) getDimensions(){
	fmt.Print("Enter no of rows:")
	fmt.Scan(&mat.row)
	fmt.Print("Enter no of col:")
	fmt.Scan(&mat.column)
}

func (mat *Matrix) getValues(){
	for i := 0; i < int(mat.row); i++ {
		for j := 0; j < int(mat.column); j++ {
			input(fmt.Sprintf("Enter value [%d][%d]",i,j),&mat.value[i][j])
		}	
	}
}
