//go:build ignore
// +build ignore

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run generate.go <01>")
		return
	}

	day := os.Args[1]

	goContent := fmt.Sprintf(`package code

func %s(input *[]string) uint32 {
    return 0
}

func %s(input *[]string) uint32 {
    return 0
}
`, "Day"+day+"Part1", "Day"+day+"Part2")

	mainContent := fmt.Sprintf(`package main

import (
	"log"

	"github.com/ChaotenHG/aoc-24/code"
	"github.com/ChaotenHG/aoc-24/utils"
)

func main() {

	input := utils.ReadLine("./input/day%s")
	input_sample := utils.ReadLine("./input_sample/day%s")

	log.Println("[Day%sPart1] Result for sample input:\033[92m", code.Day%sPart1(&input_sample), "\033[0m")
	log.Println("[Day%sPart1] Result for real input:\033[92m", code.Day%sPart1(&input), "\033[0m")

	log.Println("[Day%sPart2] Result for sample input:\033[92m", code.Day%sPart2(&input_sample), "\033[0m")
	log.Println("[Day%sPart2] Result for real input:\033[92m", code.Day%sPart2(&input), "\033[0m")

}`, day, day, day, day, day, day, day, day, day, day)

	if _, err := os.Stat("./code/day" + day + ".go"); errors.Is(err, os.ErrNotExist) {
		writeFile("./code/day"+day+".go", goContent)

		names := []string{"./input/day" + day, "./input_sample/day" + day}
		for _, v := range names {
			writeFile(v, "")
		}
	}

	writeFile("./main.go", mainContent)

	fmt.Printf("File day%s.go created successfully\n", day)
}

func writeFile(name string, content string) {
	err := os.WriteFile(name, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating Go file: %v\n", err)
		return
	}
}
