package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Position struct {
	x, y int
}

func (pos *Position) neighbors() []Position {
	neighbors := [4]Position{
		Position{x: pos.x + 1, y: pos.y},
		Position{x: pos.x - 1, y: pos.y},
		Position{x: pos.x, y: pos.y + 1},
		Position{x: pos.x, y: pos.y - 1},
	}

	return neighbors[:]
}

func (pos *Position) inBounds(terrain [][]byte) bool {
	return 0 <= pos.y && pos.y < len(terrain) && 0 <= pos.x && pos.x < len(terrain[pos.y])
}

type Node struct {
	Distance int
	InQueue  bool
	Visited  bool
}

func shortestPath(terrain [][]byte, startPos Position, target byte) int {
	nodes := make([][]Node, len(terrain))
	for y, row := range terrain {
		nodes[y] = make([]Node, len(row))

		// for x := range nodes[y] {
		// 	nodes[y][x] = visitedNode{visited: false}
		// }
	}

	nodes[startPos.y][startPos.x] = Node{Distance: 0, InQueue: true}
	queue := []Position{startPos}

	// fmt.Println(nodes)

	shortestDistance := math.MaxInt

	for len(queue) > 0 {
		pos := queue[0]
		node := &nodes[pos.y][pos.x]

		if terrain[pos.y][pos.x] == target && node.Distance < shortestDistance {
			shortestDistance = node.Distance
		}

		for _, neighbor := range pos.neighbors() {
			if neighbor.inBounds(terrain) && terrain[pos.y][pos.x] <= terrain[neighbor.y][neighbor.x]+1 {
				neighborNode := &nodes[neighbor.y][neighbor.x]

				if !neighborNode.Visited {
					distance := node.Distance + 1

					if neighborNode.InQueue {
						if distance < neighborNode.Distance {
							neighborNode.Distance = distance
						}
					} else {
						queue = append(queue, neighbor)
						neighborNode.Distance = distance
						neighborNode.InQueue = true
					}
				}

				// fmt.Println(neighbor, neighborNode.Distance)
			}
		}

		queue = queue[1:]
	}

	// for _, row := range nodes {
	// 	for _, node := range row {
	// 		fmt.Printf("%3d ", node.Distance)
	// 	}
	// 	fmt.Println()
	// }

	// return nodes[endPos.y][endPos.x].Distance
	return shortestDistance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	terrain := [][]byte{}
	// fmt.Println(cap(terrain))

	for scanner.Scan() {
		line := make([]byte, len(scanner.Bytes()))
		copy(line, scanner.Bytes())

		terrain = append(terrain, line)

		// fmt.Println(cap(terrain), len(terrain[0]))
	}

	// startPos := Position{}
	endPos := Position{}
	for r, row := range terrain {
		c := bytes.IndexByte(row, 'S')

		if c >= 0 {
			row[c] = 'a'
			// startPos = Position{y: r, x: c}
		}

		c = bytes.IndexByte(row, 'E')

		if c >= 0 {
			row[c] = 'z'
			endPos = Position{y: r, x: c}
		}
	}

	// fmt.Println(startPos, endPos)

	// fmt.Println(terrain)

	// distance := shortestPath(terrain, startPos, endPos)
	distance := shortestPath(terrain, endPos, 'a')
	fmt.Printf("Shortest Distance: %d\n", distance)
}
