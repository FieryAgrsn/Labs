package main	

import (
"testing"
)

func TestHello(t *testing.T){
 	var board [][]int
	//var check_board [][]bool
	b := grid{}
	height := 4
	width := 4
	b.check_board = make([][]bool, height+2)
	for i := 0; i < 6; i++{
		var tmp []int
		var btmp []bool
		b.check_board[i] = make([]bool, width+2)
		for j := 0; j < 6; j++{
			if i == 0 || i == height+1 || j == 0 || j == width +1{
				tmp = append(tmp, -10)
				btmp = append(btmp, true)
			}else{
				tmp = append(tmp, 0)
				btmp = append(btmp, false)
			}
			
		}
		board = append(board, tmp)
		
	}
	b.check_board[2][2] = true
	b.check_board[2][3] = true
	board[2][2] = 3
	board[2][3] = 2
	result:= printField(4, 4, board, b.check_board)
	var expected [6][6]string
	//expected = {{"|XX|","|XX|","|XX|","|XX|","|XX|","|XX|"},{"|XX|","|  |","|  |","|  |","|  |","|XX|"},{"|XX|","| 3|","| 2|","|  |","|  |","|XX|"},{"|XX|","|  |","|  |","|  |","|  |","|XX|"},{"|XX|","|  |","|  |","|  |","|  |","|XX|"},{"|XX|","|XX|","|XX|","|XX|","|XX|","|XX|"}}
	expected[0] = [6]string{"|XX|","|XX|","|XX|","|XX|","|XX|","|XX|"}
	expected[1] = [6]string{"|XX|","|  |","|  |","|  |","|  |","|XX|"}
	expected[2] = [6]string{"|XX|","|  |","| 3|","| 2|","|  |","|XX|"}
	expected[3] = [6]string{"|XX|","|  |","|  |","|  |","|  |","|XX|"}
	expected[4] = [6]string{"|XX|","|  |","|  |","|  |","|  |","|XX|"}
	expected[5] = [6]string{"|XX|","|XX|","|XX|","|XX|","|XX|","|XX|"}
	for i:=1; i<5;i++{
		for j:=1; j< 5; j++{
			
			if result[i][j]!= expected[i][j]{
				t.Error("Test failed")
			}
		}
	}
}
