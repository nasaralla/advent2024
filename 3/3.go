package main

import (
	"1/mylib"
	"fmt"
	"log"

	// "sort"
	// "strconv"
	"os"
	"regexp"
	"strings"
)

func main() {
	buf, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(buf)

	// part 1
	r, _ := regexp.Compile("mul[(][0-9]+[,][0-9]+[)]")
	all := r.FindAllString(input, -1)
	sum := 0
	for x := 0; x < len(all); x++ {
		// fmt.Println("All: ", all[x])
		firstTrim := strings.Replace(all[x], "mul(", "", 1)
		secondTrim := strings.Replace(firstTrim, ")", "", 1)
		numbers := strings.Split(secondTrim, ",")
		item1 := mylib.GetIntItem(numbers[0])
		item2 := mylib.GetIntItem(numbers[1])
		sum += item1 * item2
	}
	fmt.Println("Result 1: ", sum)

	// part 2
	r, _ = regexp.Compile("mul[(][0-9]+[,][0-9]+[)]|don't[(][)]|do[(][)]")
	all = r.FindAllString(input, -1)
	sum = 0
	canCount := true
	for x := 0; x < len(all); x++ {
		if all[x] == "don't()" {
			canCount = false
			continue
		}
		if all[x] == "do()" {
			canCount = true
			continue
		}
		if canCount {
			firstTrim := strings.Replace(all[x], "mul(", "", 1)
			secondTrim := strings.Replace(firstTrim, ")", "", 1)
			numbers := strings.Split(secondTrim, ",")
			item1 := mylib.GetIntItem(numbers[0])
			item2 := mylib.GetIntItem(numbers[1])
			sum += item1 * item2
		}
		fmt.Println(all[x])
	}
	fmt.Println("Result 2: ", sum)
}

// func indexAt(s, sep string, n int) int {
// 	idx := strings.Index(s[n:], sep)
// 	if idx > -1 {
// 		idx += n
// 	}
// 	return idx
// }
