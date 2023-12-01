package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func problem1() {
	examples, err := os.ReadFile("full-input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	examplesStr := string(examples)
	sum := 0

	for _, line := range strings.Split(examplesStr, "\n") {

		numbersMap := map[string]rune{
			"one":   '1',
			"two":   '2',
			"three": '3',
			"four":  '4',
			"five":  '5',
			"six":   '6',
			"seven": '7',
			"eight": '8',
			"nine":  '9',
		}
		var start rune
		var end rune

		for i, char := range line {
			if unicode.IsDigit(char) {
				if start == 0 {
					start = char
				}
				end = char
			} else if unicode.IsLetter(char) {
				remaining := line[i:]
				for name := range numbersMap {
					if strings.HasPrefix(remaining, name) {
						if start == 0 {
							start = numbersMap[name]
						} else if _, err := strconv.Atoi(remaining); err != nil {
							end = numbersMap[name]
						}
					}
				}
			}
		}

		combinedString := string(start) + string(end)
		combinedNumber, _ := strconv.Atoi(combinedString)
		sum += combinedNumber

	}

	fmt.Println(sum)

}

func main() {
	problem1()
}
