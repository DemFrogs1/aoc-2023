package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"slices"

	"github.com/DemFrogs1/aoc-2023/lib"
)

var instances = make(map[int]int)

func main() {
	data := lib.Parse("input.txt")

	pointsSum := 0
	for _, card := range data {
		cardId := strings.Fields(strings.Split(card, ":")[0])[1]
		c := strings.Split(strings.Split(card, ":")[1], "|")
		winningCards := strings.Fields(c[0])
		playerCards := strings.Fields(c[1])
		common, err := filterCommon(winningCards, playerCards)

		if err == nil {
			points := calcPoints(common)
			pointsSum += points
		}

		num, err := strconv.Atoi(cardId)
		if err == nil {
			addInstances(num, len(common))
		}
	}

	scratchpads := calcScratchpads()
	fmt.Println(pointsSum, scratchpads)
}

func filterCommon(arr1, arr2 []string) ([]string, error) {
	var common []string

	for _, el := range arr1 {
		if slices.Contains(arr2, el) {
			common = append(common, el)
		}
	}

	if len(common) == 0 {
		return []string{}, errors.New("no common elements found")
	}

	return common, nil
}

func calcPoints(common []string) int {
	points := 1

	for i := 1; i < len(common); i++ {
		points *= 2
	}

	return points
}
func calcScratchpads() int {
	scratchpads := 0

	for _, value := range instances {
		scratchpads += value
	}

	return scratchpads
}
func addInstances(id int, index int) {
	instances[id]++
	for i := 1; i <= index; i++ {
		instances[id+i] += 1 * instances[id]
	}
}
