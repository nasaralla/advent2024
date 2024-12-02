package mylib

import (
	"bufio"
	"log"
	"os"
)

func GetFileScanner() (*bufio.Scanner, *os.File) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner, file
}
