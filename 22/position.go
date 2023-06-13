package main

type Position struct {
	Row int
	Col int
}

type Walker struct {
	pos Position
	Dir Direction
}

func (walker Walker) Step(board Board) (Walker, bool) {
	if walker.Dir == Right || walker.Dir == Left {
		bound := board.RowBounds[walker.pos.Row]

		var newCol int

		if walker.Dir == Left {
			newCol = walker.pos.Col - 1
			if !bound.Contains(newCol) {
				newCol = bound.StopIndex
			}
		} else {
			newCol = walker.pos.Col + 1
			if !bound.Contains(newCol) {
				newCol = bound.StartIndex
			}
		}

		if board.Grid[walker.pos.Row][newCol] == Wall {
			return walker, true
		} else {
			walker.pos.Col = newCol
		}
	} else {
		bound := board.ColBounds[walker.pos.Col]

		var newRow int

		if walker.Dir == Up {
			newRow = walker.pos.Row - 1
			if !bound.Contains(newRow) {
				newRow = bound.StopIndex
			}
		} else {
			newRow = walker.pos.Row + 1
			if !bound.Contains(newRow) {
				newRow = bound.StartIndex
			}
		}

		if board.Grid[newRow][walker.pos.Col] == Wall {
			return walker, true
		} else {
			walker.pos.Row = newRow
		}
	}

	return walker, false
}

func rotateRight(pos Position, sideLength int) Position {
	return Position{
		Row: pos.Col,
		Col: sideLength - pos.Row - 1,
	}
}

func (walker Walker) StepOnCube(cube Cube) (Walker, bool) {
	newPos := walker

	if walker.Dir == Right {
		newPos.pos.Col += 1
	} else if walker.Dir == Down {
		newPos.pos.Row += 1
	} else if walker.Dir == Left {
		newPos.pos.Col -= 1
	} else if walker.Dir == Up {
		newPos.pos.Row -= 1
	}

	if !cube.board.Contains(newPos.pos) {
		faceIdx := cube.GetFaceIdx(walker.pos)
		face := cube.faces[faceIdx]

		newFaceIdx := face.neighbors[walker.Dir]
		newFace := cube.faces[newFaceIdx]

		// fmt.Printf("%p %p\n", face, newFace)

		newFaceDir, err := newFace.GetNeighborDirection(faceIdx)

		if err != nil {
			panic(err)
		}

		newDir := Direction(Mod(int(newFaceDir+2), 4))

		relPos := Position{
			Row: Mod(newPos.pos.Row-face.pos.Row, cube.sideLength),
			Col: Mod(newPos.pos.Col-face.pos.Col, cube.sideLength),
		}

		for i := 0; i < Mod(int(newDir-walker.Dir), 4); i += 1 {
			relPos = rotateRight(relPos, cube.sideLength)
		}

		newPos = Walker{
			pos: Position{
				Row: newFace.pos.Row + relPos.Row,
				Col: newFace.pos.Col + relPos.Col,
			},
			Dir: newDir,
		}
	}

	if cube.board.Grid[newPos.pos.Row][newPos.pos.Col] == Wall {
		return walker, true
	} else {
		return newPos, false
	}
}
