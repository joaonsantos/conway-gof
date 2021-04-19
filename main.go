package main

import (
	"fmt"
	"time"
)

func newBoard(size int) [][]int {
	b := make([][]int, size)

	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}

	return b
}

func printBoard(b [][]int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(fmt.Sprintf("%v ", b[i][j]))
		}

		fmt.Println("")
	}
	fmt.Println("")
}

func outOfBounds(x, y, offsetX, offsetY, lowBound, highBound int) bool {
	outOfBoundsX := (x+offsetX) < lowBound || (x+offsetX) > highBound
	outOfBoundsY := (y+offsetY) < lowBound || (y+offsetY) > highBound

	return outOfBoundsX || outOfBoundsY
}

func isAlive(c int) bool {
	return c == 1
}

func countLiveNeighbours(b [][]int, size, x, y int) int {
	liveNeighbours := 0

	for offsetX := -1; offsetX < 2; offsetX++ {
		for offsetY := -1; offsetY < 2; offsetY++ {
			if offsetX == 0 && offsetY == 0 {
				continue
			}

			if outOfBounds(x, y, offsetX, offsetY, 0, size-1) {
				continue
			}

			if b[x+offsetX][y+offsetY] == 1 {
				liveNeighbours += 1
			}

		}
	}

	return liveNeighbours
}

// Any live cell with two or three live neighbours survives.
// Any dead cell with three live neighbours becomes a live cell.
// All other live cells die in the next generation. Similarly, all other dead cells stay dead.

func live(board, next [][]int, size int) [][]int {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			neighboursAlive := countLiveNeighbours(board, size, i, j)

			if isAlive(board[i][j]) {
				if neighboursAlive == 2 || neighboursAlive == 3 {
					next[i][j] = 1 // cell survives
				} else {
					next[i][j] = 0 // cell is now dead
				}
			} else {
				if neighboursAlive == 3 {
					next[i][j] = 1 // cell is now alive
				} else {
					next[i][j] = 0 // cell remains dead
				}
			}
		}
	}

	return next
}

func mainLoop(start [][]int, size int, speed int) {
	for {
		fmt.Print("\033[H\033[2J") // clear screen
		next := newBoard(size)
		live(start, next, size)
		printBoard(next, size)
		copy(start, next)
		time.Sleep(time.Second * time.Duration(speed))

	}
}
func main() {
	const size = 5
	const speed = 1

	board := newBoard(size)
	board[0][3] = 1
	board[0][4] = 1
	board[1][1] = 1
	board[1][0] = 1
	board[1][2] = 1
	board[2][2] = 1
	board[2][3] = 1
	board[3][4] = 1
	printBoard(board, size)

	mainLoop(board, size, speed)
}
