package mylib

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFileScanner() (*bufio.Scanner, *os.File) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner, file
}

func GetIntItem(item string) int {
	integer, err := strconv.Atoi(item)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return integer
}

func GetIntDiff(a int, b int) int {
	diff := b - a
	//not using math.Abs as wanted to keep int type
	if diff < 0 {
		diff *= -1
	}
	return diff
}
