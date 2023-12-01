package main

import (
	"fmt"
	"log"
	"strings"

	"regexp"
	"strconv"

	"github.com/DemFrogs1/aoc-2023/lib"
)

func main() {
	data := lib.Parse("input.txt")
	sum1 := 0
	sum2 := 0
	wordToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for _, str := range data {
		//part 1
		sum1 += calculateSum(str)

		//part 2
		for digit, number := range wordToDigit {
			str = strings.ReplaceAll(str, digit, digit[:1]+number+digit[1:])
		}
		sum2 += calculateSum(str)
	}
	fmt.Println(sum1, sum2)
}

func calculateSum(str string) int {
	start := 0
	end := len(str) - 1
	var firstNum, lastNum string
	reg := regexp.MustCompile(`\d`)
	for start <= end {
		if reg.MatchString(string(str[start])) {
			firstNum = string(str[start])
		}
		if reg.MatchString(string(str[end])) {
			lastNum = string(str[end])
		}

		if firstNum != "" && lastNum != "" {
			break
		}

		if firstNum == "" {
			start++
		}
		if lastNum == "" {
			end--
		}
	}
	// fmt.Println(str, firstNum, lastNum)
	num, err := strconv.Atoi(firstNum + lastNum)
	if err != nil {
		log.Println(err)
	}
	return num
}
