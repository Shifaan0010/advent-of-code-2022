package main

const (
	Open = '.'
	Wall = '#'
)

type Board struct {
	Grid      []string
	RowBounds []Bound
	ColBounds []Bound
}

func (board Board) Contains(pos Position) bool {
	return pos.Row >= 0 &&
		pos.Row < len(board.Grid) &&
		board.RowBounds[pos.Row].Contains(pos.Col)
}
