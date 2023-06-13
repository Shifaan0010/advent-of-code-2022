package main

type Mover interface {
	Move(pos Walker, board Board) Walker
	MoveOnCube(pos Walker, cube Cube) Walker
}
