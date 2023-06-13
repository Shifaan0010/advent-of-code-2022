package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Priority(letter byte) int {
	return strings.Index("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", string(letter)) + 1
}

func Intersection(set1 map[byte]bool, set2 map[byte]bool) map[byte]bool {
	intersection := map[byte]bool{}

	for key, _ := range set1 {
		if set2[key] {
			intersection[key] = true
		}
	}

	return intersection
}

func StringToMap(s string) map[byte]bool {
	set := map[byte]bool{}

	for i := 0; i < len(s); i += 1 {
		set[s[i]] = true
	}

	return set
}

func PrioritySum(line string) int {
	set1 := map[byte]bool{}
	set2 := map[byte]bool{}

	for i := 0; i < len(line); i += 1 {
		if i < len(line) / 2 {
			set1[line[i]] = true
		} else {
			set2[line[i]] = true
		}
	}

	common := Intersection(set1, set2)

	sum := 0
	for ch, _ := range common {
		sum += Priority(ch)
	}

	return sum
}

func main() {
	file, err := os.Open("3.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		sum += PrioritySum(line)
	}
	fmt.Println(sum)
}
