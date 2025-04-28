package main

import (
	"1/mylib"
	"fmt"
	"log"
	"strings"
)

func main() {
	scanner, file := mylib.GetFileScanner()

	list := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, " ")
		for x := 0; x < len(val); x++ {
			list = append(list, mylib.GetIntItem(val[x]))
		}
	}

	newlist := []int{}
	for x := 0; x <= 6; x++ {
		newlist = []int{}
		for y := 0; y < len(list); y++ {
			fmt.Println("Num", list[y])
			if list[y] == 0 {
				newlist = append(newlist, 1)
			}
		}
		copy(newlist, list)
	}

	fmt.Println("Res", list)
	fmt.Println("Res2", newlist)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
