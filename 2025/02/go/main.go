package main

import (
	"fmt"
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
	result := 0
	ranges := strings.SplitSeq(input, ",")
	for r := range ranges {
		ids := strings.Split(r, "-")
		if len(ids) != 2 {
			continue
		}

		firstID, _ := strconv.Atoi(ids[0])
		lastID, _ := strconv.Atoi(ids[1])

		for id := firstID; id <= lastID; id++ {
			strID := strconv.Itoa(id)

			// check sequence
			l := len(strID)

			// skip odd length
			if l%2 != 0 {
				continue
			}

			part1 := strID[:l/2]
			part2 := strID[l/2:]
			if part1 == part2 {
				result += id
				continue
			}
		}
	}

	fmt.Println("(1) result: ", result)
}

func part2(input string) {
	result := 0
	ranges := strings.SplitSeq(input, ",")
	for r := range ranges {
		ids := strings.Split(r, "-")
		if len(ids) != 2 {
			continue
		}

		firstID, _ := strconv.Atoi(ids[0])
		lastID, _ := strconv.Atoi(ids[1])

		for id := firstID; id <= lastID; id++ {
			strID := strconv.Itoa(id)

		innerLoop:
			for chunkSize := 1; chunkSize <= len(strID)/2; chunkSize++ {
				if len(strID)%chunkSize != 0 {
					continue
				}

				splitIDs := equalSplit(strID, chunkSize)
				if len(splitIDs) < 2 {
					continue
				}

				seqPattern := splitIDs[0]
				for _, s := range splitIDs {
					if s != seqPattern {
						continue innerLoop
					}
				}

				result += id
				break
			}

		}
	}

	fmt.Println("(2) result: ", result)

}

func equalSplit(s string, seq int) []string {
	res := []string{}
	for i := 0; i < len(s); i += seq {
		if i+seq <= len(s) {
			res = append(res, s[i:i+seq])
		} else {
			res = append(res, s[i:])
		}
	}

	return res
}
