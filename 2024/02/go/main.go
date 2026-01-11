package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, _ := os.ReadFile("../input.txt")
	input := strings.TrimSpace(string(rawInput))
	lines := strings.Split(input, "\n")

	safes := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")
		fmt.Println(levels)

		check := func(levels []string) bool {
			/// 0 = nothing
			/// 1 = increasing
			/// 2 = decreasing
			getdir := func(l1, l2 int) int {
				if l1 > l2 {
					return 1
				}
				if l1 < l2 {
					return 2
				}

				return 0
			}

			l1, _ := strconv.Atoi(levels[0])
			l2, _ := strconv.Atoi(levels[1])
			prevDir := getdir(l1, l2)
			dir := 0

			for i, level := range levels {
				if i < len(levels)-1 {
					l1, _ := strconv.Atoi(level)
					l2, _ := strconv.Atoi(levels[i+1])

					prevDir = dir
					dir = getdir(l1, l2)

					diff := math.Abs(float64(l1) - float64(l2))

					if diff > 3 || diff == 0 {
						return false
					}

					if dir != prevDir && i != 0 {
						return false
					}
				}
			}

			return true
		}

		for i := range len(levels) {
			alternateLevel := []string{}
			alternateLevel = append(alternateLevel, levels[:i]...)
			alternateLevel = append(alternateLevel, levels[i+1:]...)

			fmt.Println(i, alternateLevel)
			if check(alternateLevel) {
				safes += 1
				break
			}
		}

	}

	fmt.Println("Safe: ", safes)
}
