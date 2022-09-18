package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func get_input() string {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func main() {
	input := strings.Fields(get_input())

	prevDepth := 0
	measurementCount := 0

	for index, value := range input {
		depth, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal("failed to parse input!! " + err.Error())
		}

		if index == 0 {
			prevDepth = depth
			continue
		}

		if depth > prevDepth {
			measurementCount++
		}

		prevDepth = depth
	}

	fmt.Printf("there are %v measurements that are larger than the previous measurement\n", measurementCount)
}
