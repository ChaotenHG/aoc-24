package main

import (
	"log"

	"github.com/ChaotenHG/aoc-24/code"
	"github.com/ChaotenHG/aoc-24/utils"
)

func main() {

	input := utils.ReadLine("./input/day02")
	input_sample := utils.ReadLine("./input_sample/day02")

	log.Println("[Day02Part1] Result for sample input:\033[92m", code.Day02Part1(&input_sample), "\033[0m")
	log.Println("[Day02Part1] Result for real input:\033[92m", code.Day02Part1(&input), "\033[0m")

	log.Println("[Day02Part2] Result for sample input:\033[92m", code.Day02Part2(&input_sample), "\033[0m")
	log.Println("[Day02Part2] Result for real input:\033[92m", code.Day02Part2(&input), "\033[0m")

}
