package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var oxygenRatingValues []string
	for sc.Scan() {
		oxygenRatingValues = append(oxygenRatingValues, sc.Text())
	}

	fmt.Println(oxygenRatingValues)
	fmt.Println(rune('0') == '0')
}
