package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func scanInput() (board []string, path string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanningBoard := true

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			scanningBoard = false
			continue
		}

		if scanningBoard {
			board = append(board, line)
		} else {
			path = line
		}
	}

	return board, path
}

func parsePath(pathStr string) []Mover {
	stepsRegex := regexp.MustCompile(`\d+`)
	directionRegex := regexp.MustCompile(`[RL]`)

	steps := stepsRegex.FindAllString(pathStr, -1)
	directions := directionRegex.FindAllString(pathStr, -1)

	path := []Mover{}

	for i, step := range steps {
		n, _ := strconv.Atoi(step)
		path = append(path, Steps(n))

		if i < len(directions) {
			dir := Right
			if directions[i] == "L" {
				dir = Left
			}

			path = append(path, dir)
		}
	}

	return path
}

func Mod(a, n int) int {
	return ((a % n) + n) % n
}

func main() {
	grid, pathStr := scanInput()

	rowBounds, colBounds := calcBounds(grid)

	board := Board{
		Grid:      grid,
		RowBounds: rowBounds,
		ColBounds: colBounds,
	}

	path := parsePath(pathStr)

	fmt.Println("Part 1")

	walker := Walker{
		pos: Position{
			Row: 0,
			Col: board.RowBounds[0].StartIndex,
		},
		Dir: Right,
	}

	for _, mover := range path {
		// fmt.Println(pos)
		walker = mover.Move(walker, board)
	}

	fmt.Println(walker)

	password := 1000*(walker.pos.Row+1) + 4*(walker.pos.Col+1) + int(walker.Dir)

	fmt.Printf("Password = %d\n", password)

	fmt.Println()

	fmt.Println("Part 2")

	cube := NewCube(board)

	walker2 := Walker{
		pos: Position{
			Row: 0,
			Col: board.RowBounds[0].StartIndex,
		},
		Dir: Right,
	}

	for _, mover := range path {
		// fmt.Println(walker2)
		walker2 = mover.MoveOnCube(walker2, cube)
	}

	fmt.Println(walker2)

	password2 := 1000*(walker2.pos.Row+1) + 4*(walker2.pos.Col+1) + int(walker2.Dir)
	fmt.Printf("Password = %d\n", password2)
}
