import os
import strconv
import math

fn main() {
    lines := os.read_file("../input.txt")!.trim_space().split("\n")

    mut lefts := []int{}
    mut rights := []int{}

    for line in lines {
        ints := line.split("   ")

        lefts << strconv.atoi(ints[0])!
        rights << strconv.atoi(ints[1])!
    }

    lefts.sort()
    rights.sort()

    mut total_distance := 0
    for i, left in lefts {
        total_distance += math.abs(left - rights[i])
    }

    println("Total distance: " + total_distance.str())
}
