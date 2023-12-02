package main

import (
	"fmt"
	"io/ioutil"
	"strings"
 	"strconv"
)

func is_game_possible(draws []string, given map[string]int) bool {
    for _, draw := range draws {
	diceCounts := strings.Split(draw, ",")
        for _, diceCount := range diceCounts {
		clean := strings.TrimSpace(diceCount)
		split := strings.Split(clean, " ")
		count, _ := strconv.Atoi(split[0])
		color := split[1]
		if val, ok := given[color]; ok {
			if (val < count) {
				// maximum possible is less than the current count
				return false
			}
			// else
			continue
		}
		// else, the die doesn't even exist in possible dies
		return false
	}
    }
    return true
}

func power_of_minimum_set_of_cubes(draws []string) int {
    min_required := map[string]int{"red": 1, "green": 1, "blue": 1}
    for _, draw := range draws {
	diceCounts := strings.Split(draw, ",")
        for _, diceCount := range diceCounts {
		clean := strings.TrimSpace(diceCount)
		split := strings.Split(clean, " ")
		count, _ := strconv.Atoi(split[0])
		color := split[1]
		if val, ok := min_required[color]; ok {
			if (val < count) {
				// maximum possible is less than the current count
				min_required[color] = count
			}
			// else
			continue
		}
		// else, the die doesn't even exist in possible dies
		min_required[color] = count
	}
    }
    product := 1
    for _, count := range min_required {
	product = product * count
    }
    return product
}

func main() {
    fmt.Println("Hello, World!")
    bytesRead, _ := ioutil.ReadFile("input")
    fileContent := string(bytesRead)
    lines := strings.Split(fileContent, "\n")
    
    maxPossible :=  map[string]int{"red": 12, "green": 13, "blue": 14}

    sum := 0
    power_sum := 0
    for index, line := range lines {
	if (line == "") {
		continue
	}
	game := index + 1
	data := strings.Split(line, ":")[1]
	draws := strings.Split(data, ";")
	if (is_game_possible(draws, maxPossible)) {
		sum += game
	}
        power_sum += power_of_minimum_set_of_cubes(draws)
    }
    fmt.Println(sum)
    fmt.Println(power_sum)
}

