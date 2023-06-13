package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/prob18/position"
)

func surfaceArea(object []position.Position) int {
	area := 6 * len(object)

	for _, pos1 := range object {
		for _, pos2 := range object {
			if position.ManhattenDistance(pos1, pos2) == 1 {
				area -= 1
			}
		}
	}

	return area
}

func minMaxXYZ(object []position.Position) (minPos, maxPos position.Position) {
	if len(object) == 0 {
		return minPos, maxPos
	}

	minPos = object[0]
	maxPos = object[0]

	for _, pos := range object {
		if pos.X < minPos.X {
			minPos.X = pos.X
		}
		if pos.Y < minPos.Y {
			minPos.Y = pos.Y
		}
		if pos.Z < minPos.Z {
			minPos.Z = pos.Z
		}

		if pos.X > maxPos.X {
			maxPos.X = pos.X
		}
		if pos.Y > maxPos.Y {
			maxPos.Y = pos.Y
		}
		if pos.Z > maxPos.Z {
			maxPos.Z = pos.Z
		}
	}

	return minPos, maxPos
}

func exteriorSurfaceArea(object []position.Position) int {
	minPos, maxPos := minMaxXYZ(object)
	minPos.X, minPos.Y, minPos.Z = minPos.X-1, minPos.Y-1, minPos.Z-1
	maxPos.X, maxPos.Y, maxPos.Z = maxPos.X+1, maxPos.Y+1, maxPos.Z+1

	fmt.Println(minPos, maxPos)

	grid := map[position.Position]bool{}
	enqueued := map[position.Position]bool{}

	for _, pos := range object {
		grid[pos] = true
	}

	area := 0

	start := minPos
	queue := []position.Position{start}

	enqueued[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, neighbor := range node.Neighbors() {
			if neighbor.InRange(minPos, maxPos) {
				_, inObject := grid[neighbor]
				_, wasEnqueued := enqueued[neighbor]

				if inObject {
					area += 1
				} else if !wasEnqueued {
					queue = append(queue, neighbor)
					enqueued[neighbor] = true
				}
			}
		}

		// fmt.Println(queue)
	}

	// for x := 0; x < MaxSize; x += 1 {
	// 	// fmt.Print("\033[H\033[2J")

	// 	for y := 0; y < MaxSize; y += 1 {
	// 		for z := 0; z < MaxSize; z += 1 {
	// 			if grid[x][y][z] {
	// 				fmt.Print("#")
	// 			} else if visited[x][y][z] {
	// 				fmt.Print(".")
	// 			} else {
	// 				fmt.Print(" ")
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	// 	fmt.Printf("x = %d\n", x)

	// 	// time.Sleep(time.Second * 2)
	// }

	return area
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	droplet := []position.Position{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line, " \t\n")

		pos_strs := strings.Split(line, ",")

		pos := position.FromStringSlice(pos_strs)

		droplet = append(droplet, pos)
	}

	fmt.Printf("%d\n", exteriorSurfaceArea(droplet))
}
