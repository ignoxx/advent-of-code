package main

import "core:fmt"
import "core:math"
import "core:mem"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

main :: proc() {
	raw, _ := os.read_entire_file("../input.txt")
	input := strings.trim_space(string(raw))
	lines, err := strings.split_lines(input)

	totalSafe := 0
	for line in lines {
		levels := strings.split(line, " ")

		for _, i in levels {
			variation := [dynamic]string{}
			append(&variation, ..levels[:i])
			append(&variation, ..levels[i + 1:])

			if check(variation) {
				fmt.println(variation, "passed!")
				totalSafe += 1
				break
			}
		}

	}

	fmt.println("total: ", totalSafe)
}

// https:/www.google.com

check :: proc(levels: [dynamic]string) -> bool {
	for i := 1; i < len(levels) - 1; i += 1 {
		left, _ := strconv.parse_int(levels[i - 1])
		cur, _ := strconv.parse_int(levels[i])
		right, _ := strconv.parse_int(levels[i + 1])

		if abs(left - cur) > 3 ||
		   abs(cur - right) > 3 ||
		   abs(cur - right) == 0 ||
		   abs(left - cur) == 0 {
			return false
		}

		if left > cur && cur < right {return false}
		if left < cur && cur > right {return false}
	}

	return true
}
