package code

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/ChaotenHG/aoc-24/utils"
)

func Day01Part2(input *[]string) uint32 {

	list1 := []uint32{}
	list2 := []uint32{}

	for _, line := range *input {
		numbers := strings.Split(line, "   ")

		num1, err := strconv.ParseUint(numbers[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.ParseUint(numbers[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, uint32(num1))
		list2 = append(list2, uint32(num2))
	}

	var result uint32 = 0

	for _, num := range list1 {

		count := utils.CountOccurrences(list2, num)
		if count != 0 {
			result += count * num
		}

	}

	return result
}

func Day01Part1(input *[]string) uint32 {

	list1 := []uint32{}
	list2 := []uint32{}

	for _, line := range *input {
		numbers := strings.Split(line, "   ")

		num1, err := strconv.ParseUint(numbers[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.ParseUint(numbers[1], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, uint32(num1))
		list2 = append(list2, uint32(num2))

	}

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var result uint32

	for i := 0; i < len(list1); i++ {
		result += uint32(math.Abs(float64(list1[i]) - float64(list2[i])))
	}

	return result
}
