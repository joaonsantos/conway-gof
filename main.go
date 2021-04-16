package main

import "fmt"

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

func live(b [][]int, size int) [][]int {
	next := newBoard(size)

  for i := 0; i < size; i++ {
    for j := 0; j < size; j++ {
      neighboursAlive := countLiveNeighbours(b, size, i, j)

      if isAlive(b[i][j]) {
        if (neighboursAlive == 2 || neighboursAlive == 3) {
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

// Any live cell with two or three live neighbours survives.
// Any dead cell with three live neighbours becomes a live cell.
// All other live cells die in the next generation. Similarly, all other dead cells stay dead.


func main() {
	const size = 5

	board := newBoard(size)
	board[1][2] = 1
	board[2][2] = 1
	board[3][2] = 1
	printBoard(board, size)

	next := live(board, size)
	printBoard(next, size)
}
