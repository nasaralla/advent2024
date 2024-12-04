package main

import (
	"1/mylib"
	"fmt"
	"log"

	// "sort"
	// "strconv"
	"strings"
)

const word = "XMAS"

var cords = [...][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func main() {
	scanner, file := mylib.GetFileScanner()

	input := [140][]string{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "")
		input[y] = items
		y++
	}

	countTree := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "X" {
				countTree += getDirection(i, j, input, "X")
			}
		}
	}

	fmt.Println("Count: ", countTree)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func getDirection(i int, j int, input [140][]string, letter string) int {
	nextChar, _ := nextLetter(letter)

	// scan perimeter
	sum := 0
	for x := 0; x < len(cords); x++ {
		nexti := i + cords[x][0]
		nextj := j + cords[x][1]
		if validIndex(nexti, input) && validIndex(nextj, input) {
			if input[nexti][nextj] == nextChar {
				if dive(nexti, nextj, input, cords[x], nextChar) {
					sum++
				}
			}
		}
	}
	return sum
}

func dive(i int, j int, input [140][]string, directionCord [2]int, letter string) bool {
	nextChar, index := nextLetter(letter)
	nexti := i + directionCord[0]
	nextj := j + directionCord[1]
	if !validIndex(nexti, input) || !validIndex(nextj, input) {
		return false
	}
	if input[nexti][nextj] != nextChar {
		return false
	} else {
	}
	if index+1 == len(word) {
		return true
	}
	return dive(nexti, nextj, input, directionCord, nextChar)
}

func nextLetter(letter string) (string, int) {
	i := strings.Index(word, letter)
	resp := string(word[i])
	if i < len(word)-1 {
		i++
		resp = string(word[i])
	}
	return resp, i
}

func validIndex(nexti int, input [140][]string) bool {
	if nexti < 0 {
		return false
	}
	if nexti >= len(input) {
		return false
	}
	return true
}
