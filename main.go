package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	aliveCell   = " "
	deadCell    = "."
	boardWidth  = 50
	boardHeight = 50
)

var board [boardHeight][boardWidth]bool

func initBoard() {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			num := rand.Intn(2)
			if num > 0 {
				board[i][j] = false
			} else {
				board[i][j] = true
			}
		}
	}
}

func printBoard() {
	fmt.Printf("\033[0;0f")
	for i := 0; i < len(board); i++ {
		fmt.Printf("\033[2K")
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] {
				fmt.Printf("\033[0;41m")
				fmt.Print(aliveCell)
			} else {
				fmt.Print(deadCell)
			}
			fmt.Print(" ")
			fmt.Printf("\033[0m")
		}
		fmt.Println()
	}
}

func getNumLivingNeighbors(x int, y int) int {
	numLiving := 0

	if x > 0 && board[y][x-1] {
		numLiving++
	}

	if x > 0 && y > 0 && board[y-1][x-1] {
		numLiving++
	}

	if y > 0 && board[y-1][x] {
		numLiving++
	}

	if y > 0 && x < boardWidth-1 && board[y-1][x+1] {
		numLiving++
	}

	if x < boardWidth-1 && board[y][x+1] {
		numLiving++
	}

	if y < boardHeight-1 && x < boardWidth-1 && board[y+1][x+1] {
		numLiving++
	}

	if y < boardHeight-1 && board[y+1][x] {
		numLiving++
	}

	if y < boardHeight-1 && x > 0 && board[y+1][x-1] {
		numLiving++
	}

	return numLiving
}

/*
If a living cell has two or three living neighbors, it survives
If a dead cell has three live neighbors, it becomes alive
All other cells die or stay dead
*/
func evalRules() {
	// keep a copy of the board because we cannot update the real board until all rules are evaluated
	var boardCopy [boardHeight][boardWidth]bool

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			numLivingNeighbors := getNumLivingNeighbors(j, i)

			if board[i][j] && (numLivingNeighbors == 2 || numLivingNeighbors == 3) {
				boardCopy[i][j] = true
			} else if !board[i][j] && numLivingNeighbors == 3 {
				boardCopy[i][j] = true
			} else {
				boardCopy[i][j] = false
			}
		}
	}

	// copy the updated values back to the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = boardCopy[i][j]
		}
	}
}

func main() {
	// clears the screen
	fmt.Printf("\033[H\033[2J")
	initBoard()

	for true {
		printBoard()
		evalRules()
		time.Sleep(75 * time.Millisecond)
	}
}
