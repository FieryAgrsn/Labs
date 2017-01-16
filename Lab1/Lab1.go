package main

import (
    	"fmt"
	"math/rand"
    	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	height  := 1
    	width := 1
    	mines := 1    
    	fmt.Println("Type in next:")
    	fmt.Println("x y z")
    	fmt.Println("where x = height of board, y = width of board, z = amount of mines!")
    	fmt.Scanf("%d %d %d", &height, &width, &mines)
	var board [][]int
	//var check_board [][]bool
	b := grid{}
	b.check_board = make([][]bool, height+2)
	for i := 0; i < height+2; i++{
		var tmp []int
		var btmp []bool
		b.check_board[i] = make([]bool, width+2)
		for j := 0; j < width+2; j++{
			if i == 0 || i == height+1 || j == 0 || j == width +1{
				tmp = append(tmp, -10)
				btmp = append(btmp, true)
			}else{
				tmp = append(tmp, 0)
				btmp = append(btmp, false)
				fmt.Print("O")
			}
			
		}
		board = append(board, tmp)
		
		fmt.Print("\n")
	}
	
    	if mines > height * width - 1 {
        	fmt.Println("number of mines cannot exceed number of cells")
        	return 
    	}
	fmt.Println("Keyboard config: type x and y")
	fmt.Println("where x = column, y = row")
	x :=0
	y :=0
	for{
		fmt.Scanf("%d %d", &x, &y)
		if x <= 0 || x >height || y <=0 || y > width {
			fmt.Println("Invalid input!")
		}else{
			break
		}	
	}
	b.check_board[x][y] = true
	cnt := 0
	for{
		if cnt == mines{
			break
		}else{
			for{
				x_coor := numRand(1, height+1)
				y_coor := numRand(1, width+1)
				if x_coor != x && y_coor != y && board[x_coor][y_coor] != -2{
					board[x_coor][y_coor] = -2
					cnt++
					break
				}
			}
		}
	}
	for i:=1; i<=height; i++{
		for j:=1; j<= width; j++{
			if board[i][j] == -2{
			}else{
				if board[i-1][j-1] == -2{
					board[i][j]++
				}
				if board[i-1][j] == -2{
					board[i][j]++
				} 
				if board[i-1][j+1] == -2{
					board[i][j]++
				} 
				if board[i][j+1] == -2{
					board[i][j]++
				} 
				if board[i+1][j+1] == -2{
					board[i][j]++
				} 
				if board[i+1][j] == -2{
					board[i][j]++
				} 
				if board[i+1][j-1] == -2{
					board[i][j]++
				} 
				if board[i][j-1] == -2{
					board[i][j]++
				} 
			}
		}
	}
	b.calculateField(x, y, board, height, width)
	// first print (can`t lose)
	for i:=1; i<=height; i++{
		for j:=1; j<=width; j++{	
			if b.check_board[i][j] == true {
				if board[i][j] == 0{
					fmt.Printf("| O|")
				}else{
					fmt.Printf("|%2d|", board[i][j])
				}
			}else{
				fmt.Printf("|  |")
			}
		}
		fmt.Println("\n")
	}
	// game loop
	b.gameLoop(x, y, height, width, mines, board)

	
}
func (b *grid) gameLoop( x, y, height, width, mines int, board [][]int){
	for{
		mines_cnt :=0
		for i:=1; i<=height; i++{
			for j:=1; j<= width; j++{
				if b.check_board[i][j] == false{
					mines_cnt++
				}
			}
		}
		if mines_cnt == mines{
			fmt.Println("Congratulations!!! You won!")
			break		
		}
		for{
			fmt.Scanf("%d %d", &x, &y)
			if x <= 0 || x >height || y <=0 || y > width || b.check_board[x][y] == true{
				fmt.Println("Invalid input!")
			}else{
				break
			}	
		}
		b.check_board[x][y] = true
		if board[x][y] == -2{
			fmt.Println("Game over, it was a mine there!")
			break
		}
		
		
		
		b.calculateField(x, y, board, height, width)
		printField(height, width, board, b.check_board)
		
	}
}
func printField(h, w int, board [][]int, check_board [][]bool) [][]string{
	var output [][]string
	for i := 0; i < h+2; i++{
		var tmp []string
		for j := 0; j < w+2; j++{
			if i == 0 || i == h+1 || j == 0 || j == w+1{
				tmp = append(tmp, "|XX|")
			}else{
				tmp = append(tmp, "")
			}
			
		}
		output= append(output, tmp)
		
	}
	for i:=1; i<=h; i++{
		for j:=1; j<= w; j++{
			if check_board[i][j] == true{
				if board[i][j] == 0{
					fmt.Printf("| O|")
					output[i][j] = "| O|"
				}else{
					fmt.Printf("|%2d|", board[i][j])
					output[i][j] = fmt.Sprintf("|%2d|", board[i][j])
				}
			}else{
				fmt.Printf("|  |")
				output[i][j] = "|  |"
			}
		}
		fmt.Print("\n")
	}
	return output
}
type grid struct {
    check_board [][]bool
}
 func numRand(min, max int) int {
         rand.Seed(time.Now().UTC().UnixNano())
         return rand.Intn(max-min) + min
 }
func (b *grid) calculateField(i , j int, board [][]int, h, w int){
		if board[i-1][j-1] != -2 && b.check_board[i-1][j-1] == false && i-1!=0 && j-1 !=0 {
			b.check_board[i-1][j-1] = true
			//b.calculateField(i-1, j-1, board, h, w)
		}
		if board[i-1][j] != -2 && b.check_board[i-1][j] == false && i-1!=0 {
			b.check_board[i-1][j] = true
			//b.calculateField(i-1, j, board, h, w)
		} 
		if board[i-1][j+1] != -2 && b.check_board[i-1][j+1] == false && i-1!=0 && j+1 != w+1 {
			b.check_board[i-1][j+1] = true
			//b.calculateField(i-1, j+1, board, h, w)
		} 
		if board[i][j+1] != -2 && b.check_board[i][j+1] == false && j+1 !=w+1 {
			b.check_board[i][j+1] = true
			//b.calculateField(i, j+1, board, h, w)
		} 
		if board[i+1][j+1] != -2 && b.check_board[i+1][j+1] == false && i+1!=h+1 && j+1 != w+1 {
			b.check_board[i+1][j+1] = true
			//b.calculateField(i+1, j+1, board, h, w)
		} 
		if board[i+1][j] != -2 && b.check_board[i+1][j] == false && i+1!=h+1 {
			b.check_board[i+1][j] = true
			//b.calculateField(i+1, j, board, h, w)
		} 
		if board[i+1][j-1] != -2 && b.check_board[i+1][j-1] == false && i+1!=h+1 && j-1 !=0 {
			b.check_board[i+1][j-1] = true
			//b.calculateField(i+1, j-1, board, h, w)
		} 
		if board[i][j-1] != -2 && b.check_board[i][j-1] == false &&  j-1 !=0 {
			b.check_board[i][j-1] = true
			//b.calculateField(i, j-1, board, h, w)
		}
}
