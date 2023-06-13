package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Shape int

const (
	Rock    Shape = 1
	Paper   Shape = 2
	Scissor Shape = 3
)

const (
	Win  int = 6
	Draw int = 3
	Lose int = 0
)

func CalcScore(opponentShape string, desiredOutcome string) int {
	beats := map[Shape]Shape{
		Rock:    Paper,
		Paper:   Scissor,
		Scissor: Rock,
	}

	opponentShapes := map[string]Shape{
		"A": Rock,
		"B": Paper,
		"C": Scissor,
	}

	outcomes := map[string]int{
		"X": Lose,
		"Y": Draw,
		"Z": Win,
	}

	opponent := opponentShapes[opponentShape]
	outcome := outcomes[desiredOutcome]

	score := int(outcome)

	if outcome == Win {
		score += int(beats[opponent])
	} else if outcome == Draw {
		score += int(opponent)
	} else {
		score += int(beats[beats[opponent]])
	}

	return score
}

func main() {
	file, err := os.Open("2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// fmt.Println(file.Stat())

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		slice := strings.Split(scanner.Text(), " ")

		opponent, player := slice[0], slice[1]

		score += CalcScore(opponent, player)
	}

	fmt.Println(score)
}
