package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func sign(n int) int {
	return n / abs(n)
}

func mod(a, m int) int {
	return ((a % m) + m) % m
}

func mix(numbers []int) []int {
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)

	mixed := make([]bool, len(numbers))

	for i := 0; i < len(numbers); {
		if mixed[i] {
			i += 1
			continue
		}

		// fmt.Println(i, numbers, mixed)

		number := numbersCopy[i]

		offset := mod(number, len(numbers)-1)

		for j := 0; j < offset; j += 1 {
			numbersCopy[mod(i+j, len(numbers))] = numbersCopy[mod(i+j+1, len(numbers))]
			mixed[mod(i+j, len(numbers))] = mixed[mod(i+j+1, len(numbers))]
		}

		numbersCopy[mod(i+offset, len(numbers))] = number
		mixed[mod(i+offset, len(numbers))] = true
	}

	return numbersCopy
}

func mixNRounds(numbers []int, rounds int) []int {
	type Node struct {
		number        int
		originalIndex int
	}

	currentIndexes := make([]int, len(numbers))
	for i := range currentIndexes {
		currentIndexes[i] = i
	}

	nodes := make([]Node, len(numbers))
	for i := range nodes {
		nodes[i] = Node{number: numbers[i], originalIndex: i}
	}

	for r := 0; r < rounds; r += 1 {
		for _, i := range currentIndexes {
			node := nodes[i]

			offset := mod(node.number, len(numbers)-1)

			for j := 0; j < offset; j += 1 {
				srcIndex, destIndex := mod(i+j+1, len(numbers)), mod(i+j, len(numbers))

				currentIndexes[nodes[srcIndex].originalIndex] = destIndex
				nodes[destIndex] = nodes[srcIndex]
			}

			movedIndex := mod(i+offset, len(numbers))

			currentIndexes[node.originalIndex] = movedIndex
			nodes[movedIndex] = node
		}
	}

	mixedNumbers := make([]int, len(numbers))
	for i, node := range nodes {
		mixedNumbers[i] = node.number
	}

	return mixedNumbers
}

func mixNRounds2(numbers []int, rounds int) []int {
	type Node struct {
		number int
		next   *Node
		prev   *Node
	}

	nodes := make([]*Node, len(numbers))
	for i := range nodes {
		nodes[i] = &Node{number: numbers[i]}
	}
	for i := range nodes {
		nodes[i].next = nodes[mod(i+1, len(numbers))]
		nodes[i].prev = nodes[mod(i-1, len(numbers))]
	}

	for r := 0; r < rounds; r += 1 {
		for _, node := range nodes {
			offset := mod(node.number, len(numbers)-1)

			// remove from linked list
			node.next.prev, node.prev.next = node.prev, node.next

			prevNode := node.prev

			for i := 0; i < offset; i += 1 {
				prevNode = prevNode.next
			}

			nextNode := prevNode.next

			// add to linked list
			node.next, node.prev = nextNode, prevNode
			prevNode.next, nextNode.prev = node, node
		}
	}

	mixedNumbers := make([]int, len(numbers))
	node := nodes[0]
	for i := 0; i < len(numbers); i += 1 {
		mixedNumbers[i] = node.number
		node = node.next
	}

	return mixedNumbers
}

func mixNRounds3(numbers []int, rounds int) []int {
	type Node struct {
		number int
		next   *Node
		prev   *Node
	}

	nodes := make([]Node, len(numbers))
	for i := range nodes {
		nodes[i] = Node{number: numbers[i]}
	}
	for i := range nodes {
		nodes[i].next = &nodes[mod(i+1, len(numbers))]
		nodes[i].prev = &nodes[mod(i-1, len(numbers))]
	}

	for r := 0; r < rounds; r += 1 {
		for i := range nodes {
			node := &nodes[i]
			
			offset := mod(node.number, len(numbers)-1)

			// remove from linked list
			node.next.prev, node.prev.next = node.prev, node.next

			prevNode := node.prev

			for i := 0; i < offset; i += 1 {
				prevNode = prevNode.next
			}

			nextNode := prevNode.next

			// add to linked list
			node.next, node.prev = nextNode, prevNode
			prevNode.next, nextNode.prev = node, node
		}
	}

	mixedNumbers := make([]int, len(numbers))
	node := &nodes[0]
	for i := 0; i < len(numbers); i += 1 {
		mixedNumbers[i] = node.number
		node = node.next
	}

	return mixedNumbers
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numbers := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line)

		numbers = append(numbers, num)
	}

	fmt.Println("Part 1")

	mixed := mixNRounds2(numbers, 1)

	index := slices.IndexFunc[int](mixed, func(elem int) bool { return elem == 0 })
	a, b, c := mixed[mod(index+1000, len(mixed))], mixed[mod(index+2000, len(mixed))], mixed[mod(index+3000, len(mixed))]

	fmt.Println(a, b, c)
	fmt.Printf("Sum = %d\n", a+b+c)

	// fmt.Println(mixed)

	fmt.Println()

	fmt.Println("Part 2")

	key := 811589153
	for i := range numbers {
		numbers[i] *= key
	}

	mixed2 := mixNRounds3(numbers, 10)
	index2 := slices.IndexFunc[int](mixed2, func(elem int) bool { return elem == 0 })
	a2, b2, c2 := mixed2[mod(index2+1000, len(mixed2))], mixed2[mod(index2+2000, len(mixed2))], mixed2[mod(index2+3000, len(mixed2))]

	fmt.Println(a2, b2, c2)
	fmt.Printf("Sum = %d\n", a2+b2+c2)
}
