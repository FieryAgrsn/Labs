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
	var check_board [][]bool
	for i := 0; i < height+2; i++{
		var tmp []int;
		var btmp []bool
		for j := 0; j < width+2; j++{
			if i == 0 || i == height+1 || j == 0 || j == width +1{
				tmp = append(tmp, -10)
				btmp = append(btmp, false)
			}else{
				tmp = append(tmp, 0)
				btmp = append(btmp, false)
				fmt.Print("O")
			}
		}
		board = append(board, tmp)
		check_board = append(check_board, btmp)
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
	check_board[x][y] = true
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
	// first print (can`t lose)
	for i:=1; i<=height; i++{
		for j:=1; j<=width; j++{
			if i==x && j==y{
				fmt.Println("ok")
				if board[i-1][j-1] != -2{
					if board[i-1][j-1] == 0{
						//fmt.Printf("|  |")
						check_board[i-1][j-1] = true
					}
							
				}else{
					//fmt.Printf("|  |")
				}
				if board[i-1][j] != -2{
					if board[i-1][j] == 0{
						//fmt.Printf("|  |")
						check_board[i-1][j] = true
					}
				}else{
					//fmt.Printf("|  |")
				}
				if board[i-1][j+1] != -2{
					if board[i-1][j+1] == 0{
						//fmt.Printf("|  |")
						check_board[i-1][j+1] = true
					}
				}else{
					//fmt.Printf("|  |")
				} 
				if board[i][j+1] != -2{
					if board[i][j+1] == 0{
						//fmt.Printf("|  |")
						check_board[i][j+1] = true
					}
				}else{
					//fmt.Printf("|  |")
				} 
				if board[i+1][j+1] != -2{
					if board[i+1][j+1] == 0{
						//fmt.Printf("|  |")
						check_board[i+1][j+1] = true
					}
				}else{
					//fmt.Printf("|  |")
				} 
				if board[i+1][j] != -2{
					if board[i+1][j] == 0{
						//fmt.Printf("|  |")
						check_board[i+1][j] = true
					}
				}else{
					//fmt.Printf("|  |")
				} 
				if board[i+1][j-1] != -2{
					if board[i+1][j-1] == 0{
						//fmt.Printf("|  |")
						check_board[i+1][j-1] = true
					}
				} else{
					//fmt.Printf("|  |")
				}
				if board[i][j-1] != -2{
					if board[i][j-1] == 0{
						//fmt.Printf("|  |")
						check_board[i][j-1] = true
					}
				}else{
					//fmt.Printf("|  |")
				}
			}
			if check_board[i][j] == true {
				if board[i][j] == 0{
					fmt.Printf("|  |")
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
	for{

		for{
			fmt.Scanf("%d %d", &x, &y)
			if x <= 0 || x >height || y <=0 || y > width || check_board[x][y] == true{
				fmt.Println("Invalid input!")
			}else{
				break
			}	
		}
		check_board[x][y] = true
		if board[x][y] == -2{
			fmt.Println("Game over, it was a mine there!")
			break
		}
		mines_cnt :=0
		for i:=1; i<=height; i++{
			for j:=1; j<= width; j++{
				if check_board[i][j] == false{
					mines_cnt++
				}
			}
		}
		if mines_cnt == mines + 2*height*width{
			fmt.Println("Congratulations!!! You won!")
			break		
		}
		calculateField(x, y, board, check_board)
		for i:=1; i<=height; i++{
			for j:=1; j<= width; j++{
				if check_board[i][j] == true{
					if board[i][j] == 0{
						fmt.Printf("|  |")
					}else{
						fmt.Printf("|%2d|", board[i][j])
					}

				}else{
					fmt.Printf("|  |")
				}
			}
			fmt.Print("\n")
		}
	}
	
}
 func numRand(min, max int) int {
         rand.Seed(time.Now().UTC().UnixNano())
         return rand.Intn(max-min) + min
 }
func calculateField(i, j int, board [][]int, check_board [][]bool){
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
	if board[i-1][j-1] == 0{
		check_board[i-1][j-1] = true
		calculateField(i-1, j-1, board, check_board)
	}
	if board[i-1][j] == 0{
		check_board[i-1][j] = true
		calculateField(i-1, j, board, check_board)
	} 
	if board[i-1][j+1] == 0{
		check_board[i-1][j+1] = true
		calculateField(i-1, j+1, board, check_board)
	} 
	if board[i][j+1] == 0{
		check_board[i][j+1] = true
		calculateField(i, j+1, board, check_board)
	} 
	if board[i+1][j+1] == 0{
		check_board[i+1][j+1] = true
		calculateField(i+1, j+1, board, check_board)
	} 
	if board[i+1][j] == 0{
		check_board[i+1][j] = true
		calculateField(i+1, j, board, check_board)
	} 
	if board[i+1][j-1] == 0{
		check_board[i+1][j-1] = true
		calculateField(i+1, j-1, board, check_board)
	} 
	if board[i][j-1] == 0{
		check_board[i][j-1] = true
		calculateField(i, j-1, board, check_board)
	}
}
//&&((i!=x+1 && j!=y+1)||(i!=x-1 && j!=y-1)||(i!=x+1 && j!=y-1)||(i!=x-1 && j!=y+1)||(i!=x && j!=y+1)||(i!=x && j!=y-1)||(i!=x+1 && j!=y)||(i!=x-1 && j!=y))
