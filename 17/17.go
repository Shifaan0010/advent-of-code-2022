package main

import (
	"fmt"
)

type Position struct {
	x int
	y int
}

type RockShape []string

var RockShapes = []RockShape{
	{"####"},
	{".#.", "###", ".#."},
	{"..#", "..#", "###"},
	{"#", "#", "#", "#"},
	{"##", "##"},
}

type Rock struct {
	pos   Position
	shape RockShape
}

const ChamberWidth = 7
const TotalRocks = 100000000

type Chamber [][ChamberWidth]byte

func (chamber Chamber) place(rock Rock) {
	for y := 0; y < len(rock.shape); y += 1 {
		for x := 0; x < len(rock.shape[y]); x += 1 {
			yPos := rock.pos.y + len(rock.shape) - y - 1
			xPos := rock.pos.x + x

			if rock.shape[y][x] == '#' {
				chamber[yPos][xPos] = rock.shape[y][x]
			}
		}
	}
}

func (chamber Chamber) display() {
	for y := len(chamber) - 1; y >= len(chamber)-1-50 && y >= 0; y -= 1 {
		fmt.Printf("%s\n", chamber[y])
	}
	fmt.Println()
}

type Direction byte

const (
	Left Direction = iota
	Right
	Down
	Up
)

func canMove(chamber Chamber, rock Rock, dir Direction) bool {
	// if (dir == Left && rock.pos.x <= 0) ||
	// 	(dir == Right && rock.pos.x+len(rock.shape[0])-1 >= ChamberWidth-1) ||
	// 	(dir == Down && rock.pos.y <= 0) {
	// 	return false
	// }

	dx := 0
	dy := 0

	if dir == Left {
		dx = -1
	} else if dir == Right {
		dx = 1
	} else if dir == Down {
		dy = -1
	} else if dir == Up {
		dy = 1
	}

	for y := 0; y < len(rock.shape); y += 1 {
		for x := 0; x < len(rock.shape[y]); x += 1 {
			chamberX := rock.pos.x + x + dx
			chamberY := rock.pos.y + len(rock.shape) - y - 1 + dy

			if chamberX < 0 || chamberX >= ChamberWidth || chamberY < 0 {
				return false
			}

			if rock.shape[y][x] == '#' && chamber[chamberY][chamberX] == '#' {
				return false
			}
		}
	}

	return true
}

func clearCopyPlaceDisplay(chamber Chamber, rock Rock) {
	// fmt.Print("\033[H\033[2J")

	// chamberCopy := Chamber{}
	// for y := 0; y < len(chamber); y += 1 {
	// 	row := [ChamberWidth]byte{}
	// 	copy(row[:], chamber[y][:])
	// 	chamberCopy = append(chamberCopy, row)
	// }
	// chamberCopy.place(rock)
	// chamberCopy.display()
	// time.Sleep(time.Second / 60)
}

func findCycle(chamber Chamber, jetPattern string) (int, int) {
	k := len(jetPattern)

	for i := k; i < len(chamber); i += k {
		for j := i - k; j >= k; j -= k {
			if chamber[i] == chamber[j] {
				l := chamber[j:i]
				r := chamber[i : i+(i-j)]
				eq := true
				for x := 0; x < len(l); x += 1 {
					if l[x] != r[x] {
						eq = false
						break
					}
				}

				if eq {
					fmt.Printf("%v (%d %d) %s\n", eq, i, j, chamber[i])
					return i, j
				}
			}
		}
	}

	return -1, -1
}

func main() {
	jetPattern := ""
	fmt.Scanln(&jetPattern)

	// fmt.Printf("%#v\n", jetPattern)

	jetIndex := 0

	chamber := Chamber{}
	top := 0
	tops := []int{top}

	for i := 0; i < TotalRocks; i += 1 {
		rock := Rock{Position{x: 2, y: top + 3}, RockShapes[i%len(RockShapes)]}

		for rock.pos.y+len(rock.shape) >= len(chamber) {
			row := [ChamberWidth]byte{}
			for i := 0; i < ChamberWidth; i += 1 {
				row[i] = '.'
			}

			chamber = append(chamber, row)
		}

		clearCopyPlaceDisplay(chamber, rock)

		for {
			if jetPattern[jetIndex%len(jetPattern)] == '>' {
				if canMove(chamber, rock, Right) {
					rock.pos.x += 1
				}
			} else {
				if canMove(chamber, rock, Left) {
					rock.pos.x -= 1
				}
			}

			jetIndex += 1

			clearCopyPlaceDisplay(chamber, rock)

			// fmt.Println(rock.pos)

			if !canMove(chamber, rock, Down) {
				break
			}

			rock.pos.y -= 1

			clearCopyPlaceDisplay(chamber, rock)
		}

		chamber.place(rock)

		if rock.pos.y+len(rock.shape) > top {
			top = rock.pos.y + len(rock.shape)
		}
		tops = append(tops, top)

		// displayChamber(chamber)
	}
	// displayChamber(chamber)
	// fmt.Printf("%#v\n", topChanges)

	fmt.Printf("Top = %d %d %d\n", top, len(chamber), tops[TotalRocks])

	t1, t2 := findCycle(chamber, jetPattern)
	r1, r2 := -1, -1
	for i, t := range tops {
		if i > 0 {
			if t >= t1 && tops[i-1] < t1 {
				r1 = i
				t1 = t
				// fmt.Println(i, t)
			}
			if t >= t2 && tops[i-1] < t2 {
				r2 = i
				t2 = t
				// fmt.Println(i, t)
			}
		}
	}

	fmt.Printf("Rock # %d - %d\n", r2, r1)
	fmt.Printf("Height %d - %d\n", t2, t1)

	dHeight := t1 - t2
	dRock := r1 - r2

	n := (1000000000000 - r2) / dRock
	rem := (1000000000000 - r2) % dRock
	height := t2 + n*dHeight + (tops[r2+rem] - t2)

	fmt.Printf("Height = %d\n", height)
}
