package code

import (
	"strings"
	"unicode"
)

func Day03Part1(input *[]string) uint32 {

	for _, line := range *input {

		expected := 'm'
		expectNumber := false

		op := ""

		for _, char := range line {

			if char != expected {
				if expectNumber {
					if !unicode.IsDigit(char) {
						op = ""
						continue
					}
				} else {
					op = ""
					continue
				}
			}

			switch char {
			case 'm':
				expected = 'u'
			case 'u':
				expected = 'l'
			case 'l':
				expected = '('
			case '(':
				expected = ''
				expectNumber = true
			case ',':
				expected = ''
				op += " "
			case ')':
				println(op)

				expected = 'm'
			}

			if unicode.IsDigit(char) {
				split := strings.Split(op, " ")

				if len(split) == 2 {
					if len(split[1]) > 3 {
						op = ""
						continue
					} else {
						expectNumber = false
						expected = ')'
					}
				} else if len(split) == 1 {
					if len(split[0]) > 3 {
						op = ""
						continue
					} else {
						expected = ','
					}
				}

				op += string(char)

			}

		}

	}

	return 0
}

func Day03Part2(input *[]string) uint32 {
	return 0
}
