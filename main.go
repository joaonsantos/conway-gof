package main

import "fmt"

func newBoard(size int) [][]int {
	b := make([][]int, size)

	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}

	b[1][2] = 1
	b[2][2] = 1
	b[3][2] = 1

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

func live(b [][]int) [][]int {
	next := make([][]int, len(b))
	copy(next, b)

	return next
}

// Any live cell with two or three live neighbours survives.
// Any dead cell with three live neighbours becomes a live cell.
// All other live cells die in the next generation. Similarly, all other dead cells stay dead.

func outOfBounds(curr, offsetX, offsetY, lowBound, highBound int) bool {
	outOfBoundsX := (curr+offsetX) < lowBound || (curr+offsetX) > highBound
	outOfBoundsY := (curr+offsetY) < lowBound || (curr+offsetY) > highBound

	return outOfBoundsX || outOfBoundsY
}

func isAlive(c int) bool {
	return c == 1
}

func countLiveNeighbours(b [][]int, size, x, y int) int {
	liveNeighbours := 0

	for offsetX := -1; offsetX < 2; offsetX++ {
		for offsetY := -1; offsetY < 2; offsetY++ {
			if offsetX == offsetY {
				continue
			}

			if outOfBounds(x, offsetX, offsetY, 0, size-1) {
				continue
			}

			if b[x+offsetX][y+offsetY] == 1 {
				liveNeighbours += 1
			}

		}
	}

	return liveNeighbours
}

func main() {
	const size = 5

	board := newBoard(size)
	printBoard(board, size)

	live := countLiveNeighbours(board, size, 2, 2)
	println(live)
}
