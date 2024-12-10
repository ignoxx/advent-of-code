package day1

import java.io.File

fun getInput(): List<String>{
    val text = File("src/main/kotlin/day1/input.txt").readText()
    return text.split("\n\n")
}

fun sumElfCalories(calories: String): Int {
    return calories.split("\n").sumOf { s -> s.toInt() }
}

fun main() {
    val rawCalories: List<String> = getInput()
    val parsedCalories = rawCalories.map { s -> sumElfCalories(s) }

    println("part 1: most carried calories ${parsedCalories.maxOf { i: Int -> i }}")
    println("part 2: top 3 most carried calories ${parsedCalories.sorted().takeLast(3).sum()}")
}
