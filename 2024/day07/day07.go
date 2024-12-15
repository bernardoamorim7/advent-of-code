package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day07() {
	input, err := os.Open("./day07/day07.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	var part1, part2 int = 0, 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(strings.TrimSpace(line), ": ")
		total, _ := strconv.Atoi(numsStr[0])
		nums := func(strs []string) []int {
			ns := make([]int, 0)
			for _, s := range strings.Split(strs[1], " ") {
				n, _ := strconv.Atoi(s)
				ns = append(ns, n)
			}
			return ns
		}(numsStr)

		if check(total, nums, false) {
			part1 += total
		}
		if check(total, nums, true) {
			part2 += total
		}
	}

	fmt.Printf("Part 1: %v\n", part1) // 1620690235709
	fmt.Printf("Part 2: %v\n", part2) // 145397611075341
}

func check(total int, nums []int, cond bool) bool {
	if len(nums) == 1 {
		return nums[0] == total
	}
	if check(total, append([]int{nums[0] + nums[1]}, nums[2:]...), cond) {
		return true
	}
	if check(total, append([]int{nums[0] * nums[1]}, nums[2:]...), cond) {
		return true
	}
	if cond {
		combined, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[0], nums[1]))
		if check(total, append([]int{combined}, nums[2:]...), cond) {
			return true
		}
	}
	return false
}
