package main

import "base:runtime"
import "core:fmt"
import "core:os"
import "core:strconv"
import "core:strings"

main :: proc() {
	fileInput, _ := os.read_entire_file_from_filename("../input.txt")
	input := strings.trim_space(string(fileInput))
	lines, _ := strings.split(input, "\n")

	part1(lines)
	part2(lines)
}

part2 :: proc(lines: []string) {
	dial := 50
	password := 0

	for line in lines {
		dir := line[0]
		dist, _ := strconv.parse_int(line[1:])
		totalRots := 0

		if dir == 'L' {
			inv := 0
			if dial > 0 do inv = (100 - dial)
			totalRots = int(abs(f64(inv + dist)) / 100)

			dial = (dial - dist) % 100
		} else {
			totalRots = int(abs(f64(dial + dist)) / 100)
			dial = (dial + dist) % 100
		}

		password += totalRots

		if dial < 0 do dial += 100
	}

	fmt.printf("(2) password is: %d\n", password)
}

part1 :: proc(lines: []string) {
	dial := 50
	password := 0

	for line in lines {
		dir := line[0]
		dist, _ := strconv.parse_int(line[1:])

		if dir == 'L' {
			dial = (dial - dist) % 100
		} else {
			dial = (dial + dist) % 100
		}

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			password += 1
		}
	}

	fmt.printf("(1) password is: %d\n", password)
}

