package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day05() {
	input, err := os.Open("./day05/day05.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	var part1 int = 0
	var part2 int = 0

	var orders = make([][2]int, 0)
	var updates = make([][]int, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			order := strings.Split(line, "|")
			x, _ := strconv.Atoi(order[0])
			y, _ := strconv.Atoi(order[1])
			orders = append(orders, [2]int{x, y})
		} else if line != "\n" {
			update := strings.Split(line, ",")
			updateInts := make([]int, len(update))
			for i, v := range update {
				updateInts[i], _ = strconv.Atoi(v)
			}
			updates = append(updates, updateInts)
		}
	}

	var ordersMap = make(map[int][]int)

	for _, o := range orders {
		ordersMap[o[1]] = append(ordersMap[o[1]], o[0])
	}

	for _, u := range updates {
		isOrdered := func(update []int) bool {
			for i, u := range update {
				if !isDisjoint(ordersMap[u], update[i:]) {
					return false
				}
			}
			return true
		}(u)

		if isOrdered {
			part1 += u[(len(u)-1)/2]
		} else {
			part2Array := make([]int, 0)

			elemLeft := make(map[int]struct{})
			for _, v := range u {
				elemLeft[v] = struct{}{}
			}

			for len(elemLeft) > 0 {
				for e := range elemLeft {
					keys := make([]int, 0, len(elemLeft))
					for k := range elemLeft {
						keys = append(keys, k)
					}

					if isDisjoint(ordersMap[e], keys) {
						part2Array = append(part2Array, e)
						delete(elemLeft, e)
						break
					}
				}
			}

			part2 += part2Array[(len(part2Array)-1)/2]
		}
	}

	fmt.Printf("Part 1: %v\n", part1) // 4872
	fmt.Printf("Part 2: %v\n", part2) // 5564
}

// isDisjoint checks if two arrays have no elements in common
func isDisjoint[T comparable](arr1, arr2 []T) bool {
	set := make(map[T]struct{})

	for _, v := range arr1 {
		set[v] = struct{}{}
	}

	for _, v := range arr2 {
		if _, exists := set[v]; exists {
			return false
		}
	}

	return true
}
