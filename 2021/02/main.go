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

	return strings.Trim(string(content), "\n")
}

type position struct {
	horizontal int
	depth      int
}

func (p *position) multiply() int {
	return p.horizontal * p.depth
}

func main() {
	submarinePosition := position{horizontal: 0, depth: 0}

	input := strings.Split(get_input(), "\n")

	for _, command := range input {
		parsedCmd := strings.Split(command, " ")
		cmdType := parsedCmd[0]
		cmdValue, err := strconv.Atoi(parsedCmd[1])

		if err != nil {
			log.Fatalf("what happendddd %v", err.Error())
		}

		fmt.Printf("%v %v => ", command, submarinePosition)

		switch cmdType {
		case "forward":
			submarinePosition.horizontal += cmdValue
		case "up":
			submarinePosition.depth -= cmdValue
		case "down":
			submarinePosition.depth += cmdValue
		}

		fmt.Println(submarinePosition)
	}

	fmt.Printf("submarine pos: %v, answer: %v\n", submarinePosition, submarinePosition.multiply())

}
