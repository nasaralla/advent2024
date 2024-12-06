package main

import (
	"1/mylib"
	"fmt"
	"log"
	"strings"
)

func main() {
	scanner, file := mylib.GetFileScanner()

	priority := make(map[int][]int)
	inputRule := true
	validInput := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			inputRule = false
		} else if inputRule {
			items := strings.Split(line, "|")
			priority[mylib.GetIntItem(items[0])] = append(priority[mylib.GetIntItem(items[0])], mylib.GetIntItem(items[1]))
		} else {
			items := strings.Split(line, ",")
			valid := true
			for x := 0; x < len(items)-1; x++ {
				for y := x + 1; y < len(items); y++ {
					if find(mylib.GetIntItem(items[y]), priority[mylib.GetIntItem(items[x])]) < 0 {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
			if valid {
				validInput = append(validInput, items)
			}
		}
	}
	sum := 0
	for x := 0; x < len(validInput); x++ {
		// centerItem := validInput[0]
		// centerItem := validInput[(len(validInput)-1)/2]
		fmt.Println(mylib.GetIntItem(validInput[x][(len(validInput[x])-1)/2]))
		sum = sum + mylib.GetIntItem(validInput[x][(len(validInput[x])-1)/2])
	}
	fmt.Println("Result1: ", sum)
	//part 2 incomplete
	// sum2 := 0
	// for x := 0; x < len(validInput); x++ {
	// 	// centerItem := validInput[0]
	// 	// centerItem := validInput[(len(validInput)-1)/2]
	// 	sort.Slice(validInput[x], func(i, j int) bool {
	// 		fmt.Println(mylib.GetIntItem(validInput[x][i]), " > ", mylib.GetIntItem(validInput[x][j]))
	// 		return mylib.GetIntItem(validInput[x][i]) > mylib.GetIntItem(validInput[x][j])
	// 	})
	// 	fmt.Println(validInput[x])
	// 	fmt.Println(mylib.GetIntItem(validInput[x][(len(validInput[x])-1)/2]))
	// 	sum2 = sum2 + mylib.GetIntItem(validInput[x][(len(validInput[x])-1)/2])
	// }
	// fmt.Println("Result2: ", sum2)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func find(what interface{}, where []int) (idx int) {
	for i, v := range where {
		if v == what {
			return i
		}
	}
	return -1
}
