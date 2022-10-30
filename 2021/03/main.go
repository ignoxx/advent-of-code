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

func toStringBits(bits []int) string {
	string_bits := ""
	for _, bit := range bits {
		string_bits += strconv.Itoa(bit)
	}
	return string_bits
}

func toDecimal(bits []int) int {
	string_bits := toStringBits(bits)
	decimal, _ := strconv.ParseInt(string_bits, 2, 64)
	return int(decimal)
}

func toGamma(bits []int, total int) []int {
	var gammeBits []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, bit := range bits {
		if bit >= total-bit {
			gammeBits[i] = 1
		} else {
			gammeBits[i] = 0
		}
	}

	return gammeBits
}

func toEpsilon(bits []int) []int {
	for i, v := range bits {
		if v == 1 {
			bits[i] = 0
		} else {
			bits[i] = 1
		}
	}

	return bits
}

func toOxygen(input []string) []string {
	cursor := 0
	output := input

	for {
		output = getBitOccurenceAtCursor(output, cursor, false)
		cursor++
		if len(output) == 1 {
			break
		}
	}

	return output
}

func toCarbon(input []string) []string {
	cursor := 0
	output := input

	for {
		output = getBitOccurenceAtCursor(output, cursor, true)
		cursor++
		if len(output) == 1 {
			break
		}
	}

	return output
}

func getBitOccurenceAtCursor(input []string, cursor int, shouldBeLowest bool) []string {
	output := []string{}
	bits := 0

	// count bit occurence for each position
	for _, n := range input {
		parsed_n := strings.Split(n, "")

		if parsed_n[cursor] == "1" {
			bits++
		}
	}

	// keep all inputs depending on the highest bit
	highestBit := 0
	if bits > len(input) {
		highestBit = 1
	}

	if bits == len(input) {
		highestBit = 1
	}

	if shouldBeLowest {
		if highestBit == 1 {
			highestBit = 0
		} else {
			highestBit = 1
		}
	}

	for _, n := range input {
		parsed_n := strings.Split(n, "")

		if parsed_n[cursor] == strconv.Itoa(highestBit) {
			output = append(output, n)
		}
	}

	return output
}

func main() {
	input := strings.Split(get_input(), "\n")

	var bits []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// count bit occurence for each position
	for _, n := range input {
		parsed_n := strings.Split(n, "")

		for pi, p := range parsed_n {
			if p == "1" {
				bits[pi]++
			}
		}
	}

	// fmt.Println(bits)

	gammaBits := toGamma(bits, len(input))
	gamma := toDecimal(gammaBits)
	// fmt.Println(gammaBits, toDecimal(gammaBits))

	epsilonBits := toEpsilon(gammaBits)
	epsilon := toDecimal(epsilonBits)
	// fmt.Println(epsilonBits, toDecimal(epsilonBits))

	fmt.Println(gamma, epsilon)
	power_consumption := gamma * epsilon
	fmt.Printf("the power consumption of the submarine is %v\n", power_consumption)

	// part 2
	oxygenBits := toOxygen(input)
	var oxints []int
	for _, i := range strings.Split(oxygenBits[0], "") {
		j, _ := strconv.Atoi(i)
		oxints = append(oxints, j)
	}
	oxygen := toDecimal(oxints)
	fmt.Println(oxygenBits)

	carbonBits := toCarbon(input)
	var carbints []int
	for _, i := range strings.Split(carbonBits[0], "") {
		j, _ := strconv.Atoi(i)
		carbints = append(carbints, j)
	}
	carbon := toDecimal(carbints)
	fmt.Println(carbonBits)

	other_consumption := oxygen * carbon
	fmt.Printf("the power consumption of the submarine is %v\n", other_consumption)

}
