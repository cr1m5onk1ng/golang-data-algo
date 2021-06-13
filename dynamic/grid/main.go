package main

const rows = 18
const cols = 18

func gridTravel() int {
	var grid [rows + 1][cols + 1]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			current := grid[i][j]
			righIndex := j + 1
			downIndex := i + 1
			if righIndex < cols {
				grid[i][righIndex] += current
			}
			if downIndex < rows {
				grid[downIndex][j] += current
			}
		}
	}
	return grid[rows][cols]
}

func main() {}
