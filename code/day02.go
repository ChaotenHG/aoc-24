package code

import (
	"log"
	"strconv"
	"strings"
)

func Day02Part1(input *[]string) uint32 {

	var safe uint32

	for _, line := range *input {

		split := strings.Split(line, " ")

		var numbers []int8

		for _, char := range split {

			num, err := strconv.ParseInt(char, 10, 8)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, int8(num))
		}

		if isSafe(numbers) {
			safe += 1
		}

	}

	return safe
}

func Day02Part2(input *[]string) uint32 {
	var safe uint32

	for _, line := range *input {

		split := strings.Split(line, " ")

		var numbers []int8

		for _, char := range split {

			num, err := strconv.ParseInt(char, 10, 32)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, int8(num))
		}

		if isSafe(numbers) {
			safe++
			continue
		}

		for i := 0; i < len(numbers); i++ {
			new := append(numbers[:i:i], numbers[i+1:]...)
			if isSafe(new) {
				safe++
				break
			}
		}

	}

	return safe
}

func isSafe(report []int8) bool {

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}
