package main

// 0 based index inclusive
type Bound struct {
	StartIndex int
	StopIndex  int
	// hasWall    bool
}

func (bound Bound) Contains(index int) bool {
	return bound.StartIndex <= index && index <= bound.StopIndex
}

func calcBounds(grid []string) (rows, cols []Bound) {
	maxCols := 0
	for _, row := range grid {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	rows = make([]Bound, len(grid))
	cols = make([]Bound, maxCols)

	for r := range rows {
		bound := Bound{}

		inBounds := false
		for i, tile := range grid[r] {
			// if tile == Wall {
			// 	bound.hasWall = true
			// }

			if tile == Open || tile == Wall {
				if !inBounds {
					bound.StartIndex = i
				}

				inBounds = true
				bound.StopIndex = i
			}
		}

		rows[r] = bound
	}

	for c := range cols {
		bound := Bound{}

		inBounds := false
		for i := range grid {
			if c < len(grid[i]) {
				tile := grid[i][c]

				// if tile == Wall {
				// 	bound.hasWall = true
				// }

				if tile == Open || tile == Wall {
					if !inBounds {
						bound.StartIndex = i
					}

					inBounds = true
					bound.StopIndex = i
				}
			}
		}

		cols[c] = bound
	}

	return rows, cols
}
