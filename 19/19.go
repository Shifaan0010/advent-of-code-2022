package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Material uint8

const (
	Ore Material = iota
	Clay
	Obsidian
	Geode
)

type Resources [4]int8
type Blueprint [4]Resources

func (a Resources) Add(b Resources) Resources {
	res := Resources{}

	for i, _ := range res {
		res[i] = a[i] + b[i]
	}

	return res
}

func (a Resources) Mul(factor int8) Resources {
	res := Resources{}

	for i, _ := range res {
		res[i] = a[i] * factor
	}

	return res
}

func (res Resources) IsNegative() bool {
	return res[Ore] < 0 || res[Clay] < 0 || res[Obsidian] < 0 || res[Geode] < 0
}

func CanBuild(current, required Resources) bool {
	return !current.Add(required.Mul(-1)).IsNegative()
}

func NewBlueprint(text string) (Blueprint, error) {
	regex := regexp.MustCompile(`Blueprint (\d+):[ \n]*Each ore robot costs (\d+) ore.[ \n]*Each clay robot costs (\d+) ore.[ \n]*Each obsidian robot costs (\d+) ore and (\d+) clay.[ \n]*Each geode robot costs (\d+) ore and (\d+) obsidian.`)

	matches := regex.FindStringSubmatch(text)

	if matches == nil {
		return Blueprint{}, errors.New("Text does not match regex")
	}

	vals := []int8{}
	for _, str := range matches[1:] {
		val, _ := strconv.Atoi(str)
		vals = append(vals, int8(val))
	}

	blueprint := Blueprint{
		Ore:      Resources{Ore: vals[1]},
		Clay:     Resources{Ore: vals[2]},
		Obsidian: Resources{Ore: vals[3], Clay: vals[4]},
		Geode:    Resources{Ore: vals[5], Obsidian: vals[6]},
	}

	return blueprint, nil
}

type State struct {
	RemTime   int
	Robots    Resources
	Resources Resources
}

func NextStatesPruned(state State, blueprint Blueprint) []State {
	if state.RemTime == 0 {
		return []State{}
	}

	neighbors := []State{}

	// if we can build a geode robot, ignore all other states
	if !CanBuild(state.Resources, blueprint[Geode]) {
		neighbors = append(neighbors, State{
			RemTime:   state.RemTime - 1,
			Robots:    state.Robots,
			Resources: state.Resources.Add(state.Robots),
		})
	}

	for i := Ore; i <= Geode; i += 1 {
		// skip if we have enough robots to create next type of robot
		if i <= Obsidian {
			skip := false

			for j := i + 1; j <= Geode; j += 1 {
				if CanBuild(state.Robots, blueprint[j]) {
					skip = true
					break
				}
			}

			if skip {
				continue
			}
		}

		if CanBuild(state.Resources, blueprint[i]) {
			newResources := state.Resources.Add(blueprint[i].Mul(-1)).Add(state.Robots)

			newRobots := state.Robots
			newRobots[i] += 1

			neighbors = append(neighbors, State{
				RemTime:   state.RemTime - 1,
				Robots:    newRobots,
				Resources: newResources,
			})
		}
	}

	return neighbors
}

func MaxGeodesCache(state State, blueprint Blueprint, currMax int8, cache *map[State]int8) int8 {
	if state.RemTime == 0 {
		return state.Resources[Geode]
	}

	// Min RemTime needed to add to cache
	// Needed to prevent memory usage from getting too high
	const cacheMinRemTime = 4

	// maximum amount of geode collected if we build a geode robot every step < currMax
	if int(state.RemTime*(state.RemTime-1)/2)+int(state.Robots[Geode])*state.RemTime+int(state.Resources[Geode]) < int(currMax) {
		return 0
	}

	if state.RemTime > cacheMinRemTime {
		if cachedVal, inCache := (*cache)[state]; inCache {
			return cachedVal
		}
	}

	max := currMax

	for _, nextState := range NextStatesPruned(state, blueprint) {
		collected := MaxGeodesCache(nextState, blueprint, max, cache)
		if collected >= max {
			max = collected
		}
	}

	if state.RemTime > cacheMinRemTime {
		(*cache)[state] = max
	}

	// if max >= 9 {
	// 	fmt.Printf("MaxGeodes(%v) = %v\n", state, max)
	// }

	return max
}

func MaxGeodes(startState State, blueprint Blueprint) int8 {
	cache := map[State]int8{}

	return MaxGeodesCache(startState, blueprint, 0, &cache)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanned := strings.Builder{}

	blueprints := []Blueprint{}

	for scanner.Scan() {
		line := scanner.Text()
		scanned.WriteString(line)

		blueprint, err := NewBlueprint(scanned.String())

		if err != nil {
			continue
		}

		// fmt.Printf("%v\n", blueprint)

		blueprints = append(blueprints, blueprint)

		scanned.Reset()
	}

	// qualityLevel := 0

	type IndexGeodes struct {
		index  int
		geodes int
	}

	const TotalTimePart1 = 24
	const TotalTimePart2 = 32

	var StartRobots = Resources{Ore: 1}
	var StartResources = Resources{}

	fmt.Printf("Part 1\n")

	c1 := make(chan IndexGeodes)

	for i, blueprint := range blueprints {
		go (func(i int, start State, bp Blueprint, c chan<- IndexGeodes) {
			max := int(MaxGeodes(start, bp))

			// fmt.Printf("%v\n", max)
			c <- IndexGeodes{i, max}
		})(i, State{TotalTimePart1, StartRobots, StartResources}, blueprint, c1)
	}

	qualityLevel := 0
	for i := 0; i < len(blueprints); i += 1 {
		indexGeode := <-c1
		qualityLevel += (indexGeode.index + 1) * indexGeode.geodes
		fmt.Printf("%v\n", indexGeode)
	}
	fmt.Printf("Quality Level = %d\n", qualityLevel)

	close(c1)

	// ************************************************************************

	fmt.Printf("Part 2\n")

	length := int(math.Min(3, float64(len(blueprints))))
	c2 := make(chan IndexGeodes)

	for i, blueprint := range blueprints[:length] {
		go (func(i int, start State, bp Blueprint, c chan<- IndexGeodes) {
			max := int(MaxGeodes(start, bp))

			// fmt.Printf("%v\n", max)
			c <- IndexGeodes{i, max}
		})(i, State{TotalTimePart2, StartRobots, StartResources}, blueprint, c2)
	}

	for i := 0; i < length; i += 1 {
		indexGeode := <-c2
		fmt.Printf("%v\n", indexGeode)
	}

	close(c2)
}
