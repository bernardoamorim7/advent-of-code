package day01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day01() {
	// Open input file
	input, err := os.Open("./day01/day01.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	// Create slices
	left := make([]int, 0)
	right := make([]int, 0)
	total := make([]float64, 0)

	// Read lines
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lr := strings.Split(line, "   ")
		l, _ := strconv.Atoi(lr[0])
		r, _ := strconv.Atoi(lr[1])

		left = append(left, l)
		right = append(right, r)
	}

	// Sort the slices
	slices.Sort(left)
	slices.Sort(right)

	for i, v := range left {
		l := v
		r := right[i]
		diff := math.Abs(float64(l - r))

		total = append(total, diff)
	}

	var part1 float64 = 0
	for _, v := range total {
		part1 += v
	}

	fmt.Printf("Part 1: %v\n", int(part1)) // 3246517

	var part2 int = 0

	for _, v := range left {
		part2 += v * func() int {
			count := 0
			for _, value := range right {
				if value == v {
					count++
				}
			}
			return count
		}()
	}

	fmt.Printf("Part 2: %v\n", part2) // 29379307
}
