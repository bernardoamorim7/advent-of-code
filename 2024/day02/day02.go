package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day02() {
	input, err := os.Open("./day02/day02.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	// Safe reports
	var part1 int = 0

	// Safe reports witn dampener
	var part2 int = 0

	// Read lines
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		nums := make([]int, len(strs))

		for i, str := range strs {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Error converting string to int:\n", err)
				return
			}
			nums[i] = num
		}

		if isSafe(nums) {
			part1 += 1
			part2 += 1
		} else if isSafeWithDampener(nums) {
			part2 += 1
		}
	}

	fmt.Printf("Part 1: %v\n", part1) // 252

	fmt.Printf("Part 2: %v\n", part2) // 324
}

func isSafe(nums []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(nums); i++ {
		diff := math.Abs(float64(nums[i] - nums[i-1]))

		if diff < 1 || diff > 3 {
			return false
		}
		if nums[i] > nums[i-1] {
			isDecreasing = false
		}
		if nums[i] < nums[i-1] {
			isIncreasing = false
		}
	}

	return isDecreasing || isIncreasing
}

func isSafeWithDampener(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		damp := append([]int{}, nums[:i]...)
		damp = append(damp, nums[i+1:]...)

		if isSafe(damp) {
			return true
		}
	}

	return false
}
