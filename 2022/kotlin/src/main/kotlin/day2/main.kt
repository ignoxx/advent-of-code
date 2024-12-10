package day2

import java.io.File

val choiceScore = mapOf<String, Int>(
    "A" to 1,
    "B" to 2,
    "C" to 3,

    "X" to 1,
    "Y" to 2,
    "Z" to 3
)

fun getInput(): List<String> {
    val text = File("src/main/kotlin/day2/input.txt").readText()
    return text.split("\n")
}

fun parseLinePart1(line: String): Int {
    val (opponent, me) = line.split(" ")
    var myScore: Int = choiceScore[me]!!

    // draw
    if (choiceScore[opponent] == choiceScore[me]) {
        myScore += 3
    }

    // win
    if (me == "X" && opponent == "C" || me == "Y" && opponent == "A" || me == "Z" && opponent == "B") {
        myScore += 6
    }

    // loose, 0 points

    return myScore
}

fun getWinningChoice(s: String): String {
    return when (s) {
        "A" -> "B"
        "B" -> "C"
        else -> "A"
    }
}

fun getLoosingChoice(s: String): String {
    return when (s) {
        "A" -> "C"
        "B" -> "A"
        else -> "B"
    }
}

fun parseLinePart2(line: String): Int {
    val (opponent, outcome) = line.split(" ")
    var myScore: Int = 0

    val me = when (outcome) {
        "X" -> getLoosingChoice(opponent)
        "Z" -> {
            myScore += 6
            getWinningChoice(opponent)
        }
        else -> {
            myScore += 3
            opponent
        } // draw
    }

    myScore += choiceScore[me]!!

    return myScore
}

fun main() {
    println("part 1: score by going after the strategy ${getInput().sumOf { s -> parseLinePart1(s) }}")
    println("part 2: score by new strategy ${getInput().sumOf { s -> parseLinePart2(s) }}")
}
