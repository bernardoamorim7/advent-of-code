package day04

import (
	"bufio"
	"fmt"
	"os"
)

func Day04() {
	input, err := os.Open("./day04/day04.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	rows := len(grid)
	cols := len(grid[0])

	var part1 int = countXMAS(grid, rows, cols)
	var part2 int = countXMASXShape(grid, rows, cols)

	fmt.Printf("Part 1: %v\n", part1) // 2560
	fmt.Printf("Part 2: %v\n", part2) // 1910
}

func countXMAS(grid [][]rune, rows, cols int) int {
	var count int = 0

	var directions = [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // down-right
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // up-left
		{-1, 1},  // up-right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if func(grid [][]rune, word string, row, col, dirRow, dirCol int) bool {
					for i := 0; i < len(word); i++ {
						newr := r + i*dirRow
						newc := c + i*dirCol
						if newr < 0 ||
							newr >= len(grid) ||
							newc < 0 ||
							newc >= len(grid[0]) ||
							grid[newr][newc] != rune(word[i]) {
							return false
						}
					}
					return true
				}(grid, "XMAS", r, c, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func countXMASXShape(grid [][]rune, rows, cols int) int {
	var count int = 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] != 'A' {
				continue
			}

			// Check top-left to bottom-right diagonal
			topLeftM := grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S'
			topLeftS := grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M'

			// Check top-right to bottom-left diagonal
			topRightM := grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S'
			topRightS := grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M'

			// Count valid X patterns
			if (topLeftM || topLeftS) && (topRightM || topRightS) {
				count++
			}
		}
	}

	return count
}
