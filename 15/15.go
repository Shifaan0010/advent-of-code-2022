package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Position struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func ManhattenDistance(a, b Position) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

type SensorBeacon struct {
	sensor Position
	beacon Position
}

type Range struct {
	left  int
	right int
}

func NonOverlapping(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].left < ranges[j].left
	})

	newRanges := []Range{}

	cur := Range{left: 0, right: -2}
	for _, r := range ranges {
		if cur.left <= r.left && r.left <= cur.right+1 {
			cur.right = max(cur.right, r.right)
		} else {
			if cur.left <= cur.right {
				newRanges = append(newRanges, cur)
				// fmt.Println(newRanges)
			}

			cur = r
		}
	}

	newRanges = append(newRanges, cur)
	// fmt.Println(newRanges)

	return newRanges
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	positionsRegex := regexp.MustCompile(`Sensor at x=([+-]?\d+), y=([+-]?\d+): closest beacon is at x=([+-]?\d+), y=([+-]?\d+)`)

	positions := []SensorBeacon{}

	for scanner.Scan() {
		line := scanner.Text()

		match := positionsRegex.FindStringSubmatch(line)

		if match != nil {
			vals := []int{}
			for _, s := range match[1:] {
				n, _ := strconv.Atoi(s)
				vals = append(vals, n)
			}

			sensor := Position{x: vals[0], y: vals[1]}
			beacon := Position{x: vals[2], y: vals[3]}

			positions = append(positions, SensorBeacon{sensor: sensor, beacon: beacon})

			// fmt.Printf("%v\n", positions)
		}
	}

	for y := 0; y <= 4000000; y += 1 {
		positionRanges := []Range{}

		for _, pos := range positions {
			sensor, beacon := pos.sensor, pos.beacon

			beaconDistance := ManhattenDistance(sensor, beacon)

			yDistance := abs(y - sensor.y)

			if yDistance <= beaconDistance {
				width := beaconDistance - yDistance

				left, right := sensor.x-width, sensor.x+width

				// # if beacon[1] == y:
				// #     if left == beacon[0]:
				// #         left += 1
				// #     if right == beacon[0]:
				// #         right -= 1

				if left <= right {
					posRange := Range{left: left, right: right}
					positionRanges = append(positionRanges, posRange)
				}
			}
		}

		nonOverlappingRanges := NonOverlapping(positionRanges)

		if y%400000 == 0 {
			fmt.Printf("%d%% done\n", 100*y/4000000)
		}

		if len(nonOverlappingRanges) >= 2 {
			// # y = 2766584 (-864133, 3135799) (3135801, 4102797)
			// # 3135800 * 4000000 + 2766584 = 12543202766584
			fmt.Printf("y = %2d %v\n", y, nonOverlappingRanges)
		}
	}
}
