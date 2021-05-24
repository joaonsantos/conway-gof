package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Board [][]int

func newBoard(size int) [][]int {
	b := make(Board, size)

	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}

	return b
}

func initializeBoard(b Board, size int, seed int64) {
	// initialize random seed
	rand.Seed(seed)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			b[i][j] = rand.Intn(2) // either 0 or 1
		}
	}
}

func printBoard(b Board, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%v ", b[i][j])
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

func countLiveNeighbours(b Board, size, x, y int) int {
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

func live(board, next Board, size int) Board {
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

func mainLoop(b Board, size int, delay int) {
	for {
		fmt.Print("\033[H\033[2J") // clear screen
		next := newBoard(size)
		live(b, next, size)
		printBoard(next, size)
		copy(b, next)
		time.Sleep(time.Second * time.Duration(delay))

	}
}
func main() {
	var size int
	var delay int
	var seed int64

	flag.IntVar(&size, "size", 5, "the size of the board")
	flag.IntVar(&delay, "delay", 1, "the vizualization refresh delay")
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "the seed for initializing the board")

	flag.Parse()

	board := newBoard(size)
	initializeBoard(board, size, seed)
	printBoard(board, size)

	mainLoop(board, size, delay)
}
