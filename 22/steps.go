package main

type Steps int

func (steps Steps) Move(pos Walker, board Board) Walker {
	for i := 0; i < int(steps); i += 1 {
		newPos, blocked := pos.Step(board)

		if blocked {
			break
		} else {
			pos = newPos
		}
	}

	return pos
}

func (steps Steps) MoveOnCube(pos Walker, cube Cube) Walker {
	for i := 0; i < int(steps); i += 1 {
		newPos, blocked := pos.StepOnCube(cube)

		if blocked {
			break
		} else {
			pos = newPos
		}
	}

	return pos
}
