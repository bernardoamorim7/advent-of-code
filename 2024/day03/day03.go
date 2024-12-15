package day03

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func Day03() {
	input, err := os.Open("./day03/day03.input")
	if err != nil {
		fmt.Printf("Error opening input file: \n%v\n", err)
		return
	}
	defer input.Close()

	var part1 int = 0
	var part2 int = 0

	var regexps = map[string]*regexp.Regexp{
		"do":   regexp.MustCompile(`do\(\)`),
		"dont": regexp.MustCompile(`don't\(\)`),
		"mul":  regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`),
	}
	var do bool = true

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches for do(), don't(), and mul()
		doMatches := regexps["do"].FindAllStringIndex(line, -1)
		dontMatches := regexps["dont"].FindAllStringIndex(line, -1)
		mulMatches := regexps["mul"].FindAllStringSubmatchIndex(line, -1)

		// Combine all matches into a single slice
		allMatches := append(doMatches, dontMatches...)
		for _, m := range mulMatches {
			allMatches = append(allMatches, m[:2])
		}

		// Sort matches by their starting index
		sort.Slice(allMatches, func(i, j int) bool {
			return allMatches[i][0] < allMatches[j][0]
		})

		for _, match := range allMatches {
			substr := line[match[0]:match[1]]

			if regexps["do"].MatchString(substr) {
				do = true
			} else if regexps["dont"].MatchString(substr) {
				do = false
			} else if regexps["mul"].MatchString(substr) {
				m := regexps["mul"].FindStringSubmatch(substr)
				num1, _ := strconv.Atoi(m[1])
				num2, _ := strconv.Atoi(m[2])
				part1 += (num1 * num2)

				if do {
					part2 += (num1 * num2)
				}
			}
		}
	}

	fmt.Printf("Part 1: %v\n", part1) // 164730528

	fmt.Printf("Part 2: %v\n", part2) // 70478672
}
