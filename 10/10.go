package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func run_cycle(cycleCount *int, x int, output *strings.Builder) {
	if (*cycleCount%40)-x <= 1 && (*cycleCount%40)-x >= -1 {
		(*output).WriteString("#")
	} else {
		(*output).WriteString(".")
	}

	*cycleCount += 1

	if (*cycleCount % 40) == 0 {
		(*output).WriteString("\n")
	}

	// fmt.Printf("\tx = %d\tcycles = %d\n", x, *cycleCount)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalSignalStrength := 0

	var output strings.Builder

	x := 1
	cycleCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		tokens := strings.Fields(line)

		prevX := x

		if tokens[0] == "noop" {
			// cycleCount += 1
			run_cycle(&cycleCount, x, &output)
		} else if tokens[0] == "addx" {
			n, err := strconv.Atoi(tokens[1])

			if err != nil {
				log.Fatal("Could not parse integer")
				os.Exit(1)
			}

			// cycleCount += 2
			run_cycle(&cycleCount, x, &output)
			run_cycle(&cycleCount, x, &output)
			x += n
		}

		signalStrength := 0
		if (cycleCount-20)%40 == 0 {
			signalStrength = prevX * cycleCount
		} else if tokens[0] == "addx" && (cycleCount-20)%40 == 1 {
			signalStrength = prevX * (cycleCount - 1)
		}

		totalSignalStrength += signalStrength

		// fmt.Printf("%#v\tx = %d\tcycles = %d\tSignal Strength = %d\n", tokens, x, cycleCount, signalStrength)
	}

	fmt.Printf("%s", output.String())
	fmt.Printf("Total Signal Strength = %d\n", totalSignalStrength)
}
