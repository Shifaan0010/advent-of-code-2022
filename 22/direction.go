package main

type Direction uint8

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (dir Direction) String() string {
	if dir == Right {
		return "R"
	} else if dir == Down {
		return "D"
	} else if dir == Left {
		return "L"
	} else if dir == Up {
		return "U"
	} else {
		panic("Invalid Direction")
	}
}

func (dir Direction) Move(pos Walker, board Board) Walker {
	if dir == Right {
		pos.Dir = Direction(Mod(int(pos.Dir + 1), 4))
	} else if dir == Left {
		pos.Dir = Direction(Mod(int(pos.Dir - 1), 4))
	}

	return pos
}

func (dir Direction) MoveOnCube(pos Walker, cube Cube) Walker {
	if dir == Right {
		pos.Dir = Direction(Mod(int(pos.Dir + 1), 4))
	} else if dir == Left {
		pos.Dir = Direction(Mod(int(pos.Dir - 1), 4))
	}

	return pos
}
