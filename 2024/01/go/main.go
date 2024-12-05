package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rawInput, _ := os.ReadFile("../input.txt")
	input := strings.TrimSpace(string(rawInput))
	lines := strings.Split(input, "\n")

	leftLocationIDs := []int{}
	rightLocationIDs := []int{}
	rightIDOcc := map[int]int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		ids := strings.Split(line, "   ")

		leftID, _ := strconv.Atoi(ids[0])
		rightID, _ := strconv.Atoi(ids[1])

		rightIDOcc[rightID] += 1

		leftLocationIDs = append(leftLocationIDs, leftID)
		rightLocationIDs = append(rightLocationIDs, rightID)
	}

	sort.Slice(leftLocationIDs, func(i, j int) bool {
		return leftLocationIDs[i] < leftLocationIDs[j]
	})

	sort.Slice(rightLocationIDs, func(i, j int) bool {
		return rightLocationIDs[i] < rightLocationIDs[j]
	})

	totalDistance := float64(0)
	similarityScore := 0
	for i := range len(rightLocationIDs) {
		totalDistance += math.Abs(float64(leftLocationIDs[i] - rightLocationIDs[i]))
		similarityScore += leftLocationIDs[i] * rightIDOcc[leftLocationIDs[i]]
	}

	// Part 1
	fmt.Printf("Total distance: %f\n", totalDistance)

	// Part 2
	fmt.Printf("Similarity score: %d\n", similarityScore)

}
