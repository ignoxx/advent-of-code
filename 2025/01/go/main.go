package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileInput, _ := os.ReadFile("../input.txt")
	input := strings.TrimSpace(string(fileInput))

	part1(input)
	part2(input)
}

func part1(input string) {
	dial := 50
	password := 0
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		dir := string(line[0])
		dist, _ := strconv.Atoi(line[1:])

		if dir == "L" {
			dial = (dial - dist) % 100
		} else {
			dial = (dial + dist) % 100
		}

		if dial < 0 {
			dial += 100
		}

		if dial == 0 {
			password++
		}
	}

	fmt.Println("(1) password: ", password)
}

func part2(input string) {
	dial := 50
	password := 0

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		dir := string(line[0])
		dist, _ := strconv.Atoi(line[1:])
		totalRots := 0

		if dir == "L" {
			inverted := 0
			if dial > 0 {
				inverted = (100 - dial)
			}

			totalRots = int(math.Abs((float64(inverted + dist)) / 100))
			dial = (dial - dist) % 100
		} else {
			totalRots = int(math.Abs((float64(dial + dist)) / 100))
			dial = (dial + dist) % 100
		}

		if dial < 0 {
			dial += 100
		}

		password += totalRots
	}

	fmt.Println("(2) password: ", password)
}
