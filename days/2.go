package days

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func Two() {
	lines := utils.LoadLines("2.txt")
	total1 := 0
	total2 := 0

	for _, l := range lines {
		split := splitGameLine(l)
		id := getId(split)
		games := strings.Split(split[1], ";")
		r, g, b := handleGames(games)

		total2 += (r * b * g)

		if r <= 12 && g <= 13 && b <= 14 {
			total1 += id
		}
	}
	fmt.Printf("Day 2 Part 1: %d\n", total1)
	fmt.Printf("Day 2 Part 2: %d\n", total2)
}

func handleGames(games []string) (int, int, int) {

	r, g, b := 0, 0, 0
	for _, game := range games {
		colors := strings.Split(game, ",")
		for _, color := range colors {
			i, c := getNumColor(color)

			switch c {
			case "red":
				r = getMax(r, i)
			case "green":
				g = getMax(g, i)
			case "blue":
				b = getMax(b, i)
			}
		}
	}
	return r, g, b
}

func splitGameLine(line string) []string {
	return strings.Split(line, ":")
}

// takes output from splitGameLine as input
func getId(gameLine []string) int {
	id := strings.Split(gameLine[0], " ")[1]
	ret, _ := strconv.Atoi(id)
	return ret
}
func getNumColor(input string) (int, string) {
	split := strings.Split(input, " ")
	i, _ := strconv.Atoi(split[1])
	return i, split[2]
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
