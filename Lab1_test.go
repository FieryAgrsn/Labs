package main	

import "testing"

func TestHello(t *testing.T){
 	var board [][]int
	//var check_board [][]bool
	b := grid{}
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
	expected: = ("{|XX||XX||XX||XX||XX||XX|},{|XX||  ||  ||  ||  ||XX|},{|XX|| 3|| 2||  ||  ||XX|},{|XX||  ||  ||  ||  ||XX|},{|XX||  ||  ||  ||  ||XX|},{|XX||XX||XX||XX||XX||XX|}")
	if result != expected {
                t.Error("Test failed")
        }
	

}
