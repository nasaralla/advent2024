package main

import (
	"1/mylib"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner, file := mylib.GetFileScanner()
	var listLeft []int
	var listRight []int

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")
		listLeft = append(listLeft, getIntItem(items[0]))
		listRight = append(listRight, getIntItem(items[1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	//sort the lists
	sort.Slice(listLeft, func(i, j int) bool {
		return listLeft[i] < listLeft[j]
	})
	sort.Slice(listRight, func(i, j int) bool {
		return listRight[i] < listRight[j]
	})

	result := 0

	for x := 0; x < len(listLeft); x++ {
		diff := listRight[x] - listLeft[x]
		//not using math.Abs as wanted to keep int type
		if diff < 0 {
			diff *= -1
		}
		result += diff
	}
	fmt.Println("Sum of difference: ", result)

	//Similarity score
	score := 0
	for x := 0; x < len(listLeft)-1; x++ {
		itemLeft := listLeft[x]
		itemRight := listRight[0]
		repeat := 0
		rightIndex := 0

		for itemLeft >= itemRight {
			if itemLeft == itemRight {
				repeat++
			} else {
				repeat = 0
			}
			itemRight = listRight[rightIndex]
			rightIndex++
		}
		score = score + (itemLeft * repeat)
	}

	fmt.Println("Similarity score: ", score)
}

func getIntItem(item string) int {
	integer, err := strconv.Atoi(item)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return integer
}
