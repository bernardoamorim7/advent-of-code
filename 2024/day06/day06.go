package day06

import (
	"bufio"
	"fmt"
	"os"
)

func Day06() {
	input, err := os.Open("./day06/day06.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	var directions = map[rune][2]int{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	var nextDirections = map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	const BLOCKAGE rune = '#'

	var gridArray = make([]string, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		gridArray = append(gridArray, line)
	}

	var startingPosition [2]int
	var startingPositionSet bool = false
	var startingDirection rune = '^'
	var gridWidth int = len(gridArray[0])
	var gridHeight int = len(gridArray)

	for y, line := range gridArray {
		if !startingPositionSet {
			for x, char := range line {
				if char == '^' {
					startingPosition = [2]int{x, y}
					startingPositionSet = true
					break
				}
			}
		}
	}

	var grid = func(gridArray []string) map[[2]int]rune {
		var m = make(map[[2]int]rune)
		for y, line := range gridArray {
			for x, char := range line {
				m[[2]int{x, y}] = char
			}
		}
		return m
	}(gridArray)

	var part1 = func(grid map[[2]int]rune, startingPosition [2]int, startingDirection rune) int {
		var stepsTaken = make(map[[2]int]struct{}, 0)

		// Add starting position
		stepsTaken[startingPosition] = struct{}{}

		var currentPosition = startingPosition
		var currentDirection = startingDirection

		for {
			xPos, yPos := currentPosition[0], currentPosition[1]

			xDir, yDir := directions[currentDirection][0], directions[currentDirection][1]

			var nextPosition = [2]int{xPos + xDir, yPos + yDir}

			xNext, yNext := nextPosition[0], nextPosition[1]

			// Check boundaries
			if xNext < 0 || xNext >= gridWidth || yNext < 0 || yNext >= gridHeight {
				break
			}

			// Change direction if blockage is encountered
			if grid[nextPosition] == BLOCKAGE {
				currentDirection = nextDirections[currentDirection]
				continue
			}

			// Take a step
			currentPosition = nextPosition

			stepsTaken[nextPosition] = struct{}{}
		}

		return len(stepsTaken)
	}(grid, startingPosition, startingDirection)

	var part2 = func(grid map[[2]int]rune, startingPosition [2]int, startingDirection rune) int {
		var loops int = 0

		for pos := range grid {
			// Skip starting position and existing obstacles
			if pos == startingPosition || grid[pos] == BLOCKAGE {
				continue
			}

			// Similar to part1, but with state tracking
			stepsTaken := make(map[[3]int]struct{})
			currentPosition := startingPosition
			currentDirection := startingDirection

			for {
				state := [3]int{currentPosition[0], currentPosition[1], int(currentDirection)}

				if _, exists := stepsTaken[state]; exists {
					loops++
					break
				}
				stepsTaken[state] = struct{}{}

				nextPosition := [2]int{
					currentPosition[0] + directions[currentDirection][0],
					currentPosition[1] + directions[currentDirection][1],
				}

				// Check boundaries
				if nextPosition[0] < 0 || nextPosition[0] >= gridWidth || nextPosition[1] < 0 || nextPosition[1] >= gridHeight {
					break
				}

				// Check for obstacles (including our test obstacle)
				if grid[nextPosition] == BLOCKAGE || nextPosition == pos {
					currentDirection = nextDirections[currentDirection]
					continue
				}

				currentPosition = nextPosition
			}
		}

		return loops
	}(grid, startingPosition, startingDirection)

	fmt.Printf("Part 1: %v\n", part1) // 4988
	fmt.Printf("Part 2: %v\n", part2) // 1697
}
