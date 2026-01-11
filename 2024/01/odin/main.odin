package main

import "core:fmt"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

main :: proc() {
	raw, _ := os.read_entire_file("../input.txt")
	input := strings.trim_space(string(raw))
	lines, err := strings.split_lines(input)

	lefts := [dynamic]int{}
	rights := [dynamic]int{}


	for line in lines {
		ids := strings.split(line, "   ")
		left, _ := strconv.parse_int(ids[0])
		right, _ := strconv.parse_int(ids[1])
		append_elem(&lefts, left)
		append_elem(&rights, right)
	}

	slice.sort(lefts[:])
	slice.sort(rights[:])

	totalDistance := 0
	for v, i in lefts {
		totalDistance += abs(v - rights[i])
	}

	fmt.printf("Total distance: %d\n", totalDistance)
}

