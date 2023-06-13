package main

import "fmt"
import "errors"

type Face struct {
	pos Position
	// neighbors [4]*Face // Need to stop using pointers since its becomes invalid when cloned
	neighbors map[Direction]int
}

func (face Face) GetNeighborDirection(neighborIdx int) (Direction, error) {
	for dir, idx := range face.neighbors {
		if idx == neighborIdx {
			return dir, nil
		}
	}

	return 99, errors.New("Not a neighbor")
}

func Neighbors(cube Cube, faceIdx int, sideLength int, board Board) []Face {
	face := cube.faces[faceIdx]

	neighbors := [4]Face{
		{pos: Position{Row: face.pos.Row - sideLength, Col: face.pos.Col}, neighbors: map[Direction]int{Down: faceIdx}},
		{pos: Position{Row: face.pos.Row + sideLength, Col: face.pos.Col}, neighbors: map[Direction]int{Up: faceIdx}},
		{pos: Position{Row: face.pos.Row, Col: face.pos.Col - sideLength}, neighbors: map[Direction]int{Right: faceIdx}},
		{pos: Position{Row: face.pos.Row, Col: face.pos.Col + sideLength}, neighbors: map[Direction]int{Left: faceIdx}},
	}

	validNeighbors := []Face{}

	for _, neighbor := range neighbors {
		if board.Contains(neighbor.pos) {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}

	return validNeighbors
}

type Cube struct {
	board      Board
	faces      [6]Face
	sideLength int
}

func (cube Cube) GetFaceIdx(pos Position) int {
	pos.Row -= Mod(pos.Row, cube.sideLength)
	pos.Col -= Mod(pos.Col, cube.sideLength)

	for i := range cube.faces {
		if cube.faces[i].pos == pos {
			return i
		}
	}

	return -1
}

func NewCube(board Board) Cube {
	sideLength := calcSideLength(board)

	fmt.Printf("Side length = %d\n", sideLength)

	cube := Cube{board: board, sideLength: sideLength}

	// top leftmost face
	cube.faces[0] = Face{
		pos:       Position{Row: 0, Col: board.RowBounds[0].StartIndex},
		neighbors: map[Direction]int{},
	}

	// bfs cube faces
	length := 1
	for i := 0; i < len(cube.faces); i += 1 {
		for _, neighbor := range Neighbors(cube, i, cube.sideLength, cube.board) {
			var neighborIdx int

			alreadyAdded := false
			for j := 0; j < length; j += 1 {
				if cube.faces[j].pos == neighbor.pos {
					alreadyAdded = true
					neighborIdx = j
				}
			}

			if !alreadyAdded {
				cube.faces[length] = neighbor
				neighborIdx = length
				length += 1
			}

			// connect back to current node
			for dir, index := range cube.faces[neighborIdx].neighbors {
				if index == i {
					//
					if dir == Right {
						cube.faces[i].neighbors[Left] = neighborIdx
					} else if dir == Left {
						cube.faces[i].neighbors[Right] = neighborIdx
					} else if dir == Up {
						cube.faces[i].neighbors[Down] = neighborIdx
					} else if dir == Down {
						cube.faces[i].neighbors[Up] = neighborIdx
					}
				}
			}
		}
	}

	connectFaces(&cube)

	// fmt.Printf("%v\n", cube.faces)

	return cube
}

func connectFaces(cube *Cube) {
	// for i, face := range cube.faces {
	// 	fmt.Printf("%d: %v\n", i, face)
	// }

	// fmt.Println()

	for disconnected := true; disconnected; {
		disconnected = false

		for i := range cube.faces {
			face := cube.faces[i]

			for curDir := Right; curDir <= Up; curDir += 1 {

				nextDir := Direction(Mod(int(curDir)+1, 4))

				// fmt.Printf("curDir: %s nextDir: %s\n", curDir, nextDir)

				// check if neighbouring faces exist
				curNeighborIdx, curNeighborExists := face.neighbors[curDir]
				nextNeighborIdx, nextNeighborExists := face.neighbors[nextDir]

				// fmt.Printf("curNeighborIdx: %d nextNeighborIdx: %d\n", curNeighborIdx, nextNeighborIdx)

				if curNeighborExists && nextNeighborExists {
					curNeighborDir, err := cube.faces[curNeighborIdx].GetNeighborDirection(i)
					if err != nil {
						panic(fmt.Sprintf("%s %d %v", "Neighbors not consistent", i, curDir))
					}

					nextNeighborDir, err := cube.faces[nextNeighborIdx].GetNeighborDirection(i)
					if err != nil {
						panic(fmt.Sprintf("%s %d %v", "Neighbors not consistent", i, curDir))
					}

					// fmt.Printf("curNeighbor: %v\n", cube.faces[curNeighborIdx])
					// fmt.Printf("nextNeighbor: %v\n", cube.faces[nextNeighborIdx])
					// fmt.Printf("curNeighborDir: %s nextNeighborDir: %s\n", curNeighborDir, nextNeighborDir)

					curNeighborDir = Direction(Mod(int(curNeighborDir)-1, 4))
					nextNeighborDir = Direction(Mod(int(nextNeighborDir)+1, 4))

					// fmt.Printf("curNeighborDir: %s nextNeighborDir: %s\n", curNeighborDir, nextNeighborDir)

					// check if neighbouring faces are connected
					_, connected1 := cube.faces[nextNeighborIdx].neighbors[nextNeighborDir]
					_, connected2 := cube.faces[curNeighborIdx].neighbors[curNeighborDir]

					if !connected1 && !connected2 {
						// fmt.Printf("%d %d\n", curNeighborIdx, nextNeighborIdx)
						disconnected = true

						// connect faces
						cube.faces[curNeighborIdx].neighbors[curNeighborDir] = nextNeighborIdx
						cube.faces[nextNeighborIdx].neighbors[nextNeighborDir] = curNeighborIdx
					} else if !connected1 || !connected2 {
						panic(fmt.Sprintf("Error %d %s %s , %d %s %s", curNeighborIdx, curDir, curNeighborDir, nextNeighborIdx, nextDir, nextNeighborDir))
					}
				}
			}
		}
	}

	// fmt.Println()

	// for i, face := range cube.faces {
	// 	fmt.Printf("%d: %v\n", i, face)
	// }
	// fmt.Println()
}

func calcSideLength(board Board) int {
	minLength := -1

	for _, boundList := range [][]Bound{board.RowBounds, board.ColBounds} {
		for _, bound := range boundList {
			length := bound.StopIndex - bound.StartIndex + 1
			if minLength == -1 || length < minLength {
				minLength = length
			}
		}
	}

	return minLength
}
