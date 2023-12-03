package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DemFrogs1/aoc-2023/lib"
)

var gameThreshold = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	data := lib.Parse("input.txt")

	score := 0
	score2 := 0
	for _, g := range data {
		game := strings.Split(g, ":")
		gameId := strings.Fields(game[0])[1]
		subsets := strings.Split(strings.ReplaceAll(game[1], ";", ","), ",")

		possible := true

		set := make(map[string]int)

		for _, subset := range subsets {
			s := strings.Split(strings.TrimSpace(subset), " ")
			score, err := strconv.Atoi(s[0])
			if err != nil {
				log.Println(err)
			}
			if set[s[1]] < score {
				set[s[1]] = score
			}
			if gameThreshold[s[1]] < score {
				possible = false
			}
		}

		if possible {
			num, err := strconv.Atoi(gameId)
			if err != nil {
				log.Println(err)
			}
			score += num
		}
		score2 += calcPower(set)
	}
	fmt.Println(score, score2)
}

func calcPower(set map[string]int) int {
	power := 1
	for _, value := range set {
		power *= value
	}
	return power
}
