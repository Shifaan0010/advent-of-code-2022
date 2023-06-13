package position

import "strconv"

type Position struct {
	X int
	Y int
	Z int
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func (pos Position) Neighbors() []Position {
	return []Position{
		{pos.X + 1, pos.Y, pos.Z},
		{pos.X - 1, pos.Y, pos.Z},
		{pos.X, pos.Y + 1, pos.Z},
		{pos.X, pos.Y - 1, pos.Z},
		{pos.X, pos.Y, pos.Z + 1},
		{pos.X, pos.Y, pos.Z - 1},
	}
}

func (pos Position) InRange(start, end Position) bool {
	return start.X <= pos.X && pos.X <= end.X && start.Y <= pos.Y && pos.Y <= end.Y && start.Z <= pos.Z && pos.Z <= end.Z
}

func ManhattenDistance(pos1, pos2 Position) int {
	return abs(pos1.X-pos2.X) + abs(pos1.Y-pos2.Y) + abs(pos1.Z-pos2.Z)
}

func FromStringSlice(strs []string) Position {
	X, _ := strconv.Atoi(strs[0])
	Y, _ := strconv.Atoi(strs[1])
	Z, _ := strconv.Atoi(strs[2])

	return Position{X, Y, Z}
}
