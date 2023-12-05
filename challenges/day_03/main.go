package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/DemFrogs1/aoc-2023/lib"
)

func main() {
	data := lib.Parse("input.txt")
	part1(data)
	part2(data)

}
func part1(data []string) {
	sum := 0
	for strIndex, schematic := range data {
		reg := regexp.MustCompile(`\d+`)
		indexes := reg.FindAllIndex([]byte(schematic), -1)
		for _, index := range indexes {
			var up, down, left, right string
			start := index[0]
			end := index[len(index)-1]
			minus := 0
			plus := 0

			if start > 0 {
				left = string(schematic[start-1])
				minus = 1
			}
			if end < len(schematic)-1 {
				right = string(schematic[end])
				plus = 1
			}

			if strIndex > 0 {
				up = data[strIndex-1][start-minus : end+plus]
			}
			if strIndex < len(data)-1 {
				down = data[strIndex+1][start-minus : end+plus]
			}

			if symbolExists(up + down + left + right) {
				num, err := strconv.Atoi(schematic[start:end])
				if err == nil {
					sum += num
				}
			}
		}
	}
	fmt.Println(sum)
}
func part2(data []string) {
	sum := 0
	for strIndex, schematic := range data {
		reg := regexp.MustCompile(`\*`)
		indexes := reg.FindAllIndex([]byte(schematic), -1)
		for _, index := range indexes {
			var up, down int
			start := index[0]
			end := index[len(index)-1]
			var adjacentNumber []int

			if start > 0 {
				num, err := getAdjacentNumber(start, -1, schematic)
				if err == nil {
					adjacentNumber = append(adjacentNumber, num...)
				}
			}
			if end < len(schematic)-1 {
				num, err := getAdjacentNumber(-1, end, schematic)
				if err == nil {
					adjacentNumber = append(adjacentNumber, num...)
				}
			}

			if strIndex > 0 {
				up = strIndex - 1
				num, err := getAdjacentNumber(start, end, data[up])
				if err == nil {
					adjacentNumber = append(adjacentNumber, num...)
				}
			}
			if strIndex < len(data)-1 {
				down = strIndex + 1
				num, err := getAdjacentNumber(start, end, data[down])
				if err == nil {
					adjacentNumber = append(adjacentNumber, num...)
				}
			}
			if len(adjacentNumber) == 2 {
				sum += adjacentNumber[0] * adjacentNumber[1]
			}
		}
	}
	fmt.Println(sum)
}
func symbolExists(s string) bool {
	for _, char := range s {
		if string(char) != "." {
			return true
		}
	}
	return false
}

func checkAdjacent(startIndex int, endIndex int, digitIndex [][]int) [][]int {
	var value [][]int
	for _, row := range digitIndex {
		if startIndex >= row[0] && startIndex <= row[len(row)-1] || endIndex >= row[0] && endIndex <= row[len(row)-1] {
			value = append(value, row)
		}
	}
	return value
}

func getAdjacentNumber(startIndex int, endIndex int, str string) ([]int, error) {
	reg := regexp.MustCompile(`\d+`)

	matches := reg.FindAllIndex([]byte(str), -1)
	numberIndex := checkAdjacent(startIndex, endIndex, matches)
	if len(numberIndex) <= 0 {
		return []int{}, errors.New("no adjacent number found")
	}

	var numbers []int

	for _, i := range numberIndex {
		num, err := strconv.Atoi(str[i[0]:i[len(i)-1]])
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers, nil
}
