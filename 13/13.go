package main

import (
	"bufio"
	"os"
)

type List struct {
	isArray bool
	array   []List
	n       int
}

func inOrder(pair []string) bool {
	depths := []int{0, 0}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	pair := []string{}
	index := 1

	indexSum := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		pair = append(pair, line)

		if len(pair) == 2 {
			if inOrder(pair) {
				indexSum += index
			}

			pair = []string{}
			index += 1
		}

		// fmt.Printf("%s\n", line)
	}
}
