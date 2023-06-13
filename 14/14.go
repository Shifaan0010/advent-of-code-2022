package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x, y int
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func parsePosition(s string) Position {
	strs := strings.Split(strings.Trim(s, " "), ",")

	x, err := strconv.Atoi(strs[0])

	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(strs[1])

	if err != nil {
		log.Fatal(err)
	}

	return Position{x: x, y: y}
}

type Material byte

const (
	Air Material = iota
	Rock
	Sand
)

const Width = 400
const Start = 300

func printCrossSection(crossSection [][Width]Material) {
	for _, row := range crossSection {
		for _, m := range row {
			if m == Air {
				fmt.Print(".")
			} else if m == Rock {
				fmt.Print("#")
			} else if m == Sand {
				fmt.Print("o")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}

func getCrossSection(rocks [][]Position) [][Width]Material {
	crossSection := [][Width]Material{}

	k := 2

	for _, segment := range rocks {
		// crossSection = append(crossSection, [Width]Material{})

		for i := 0; i < len(segment)-k+1; i += 1 {
			pos1 := segment[i]
			pos2 := segment[i+1]

			length := 0

			if pos1.x != pos2.x {
				length = Abs(pos1.x - pos2.x)
			} else {
				length = Abs(pos1.y - pos2.y)
			}

			for j := 0; j <= length; j += 1 {
				pos := Position{}

				pos.x = pos1.x + (pos2.x-pos1.x)*j/length - Start
				pos.y = pos1.y + (pos2.y-pos1.y)*j/length

				for len(crossSection) <= pos.y {
					crossSection = append(crossSection, [Width]Material{})
				}

				crossSection[pos.y][pos.x] = Rock

				// fmt.Print(pos)
			}

			// fmt.Println()
		}
	}

	crossSection = append(crossSection, [Width]Material{})
	crossSection = append(crossSection, [Width]Material{})
	for i := 0; i < len(crossSection[0]); i += 1 {
		crossSection[len(crossSection) - 1][i] = Rock
	}

	// fmt.Println(crossSection)
	// printCrossSection(crossSection)

	return crossSection
}

func dropSand(crossSection [][Width]Material, startPos Position) int {
	count := 0

	sandPos := startPos
	for {
		if crossSection[sandPos.y+1][sandPos.x] == Air {
			sandPos.y += 1
		} else if crossSection[sandPos.y+1][sandPos.x-1] == Air {
			sandPos.y += 1
			sandPos.x -= 1
		} else if crossSection[sandPos.y+1][sandPos.x+1] == Air {
			sandPos.y += 1
			sandPos.x += 1
		} else {
			crossSection[sandPos.y][sandPos.x] = Sand
			count += 1

			if sandPos == startPos {
				break
			}

			sandPos = startPos
		}

		// if sandPos.y == len(crossSection) - 1 {
		// 	break
		// }
	}

	return count
}

func sandCount(rocks [][]Position) int {
	caveCrossSection := getCrossSection(rocks)

	count := dropSand(caveCrossSection, Position{500 - Start, 0})

	printCrossSection(caveCrossSection)

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rocks := [][]Position{}

	for scanner.Scan() {
		line := scanner.Text()

		positions := []Position{}

		tokens := strings.Split(line, "->")
		for _, token := range tokens {
			positions = append(positions, parsePosition(token))
		}

		rocks = append(rocks, positions)
	}

	// fmt.Printf("%v\n", rocks)

	fmt.Printf("Units of Sand: %d\n", sandCount(rocks));
}
