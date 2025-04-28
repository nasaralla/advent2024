package main

import (
	"1/mylib"
	"fmt"
	"log"
	"strings"
)

func main() {
	scanner, file := mylib.GetFileScanner()

	block := ""
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")
		index := 0
		for x := 0; x < len(values); x++ {
			if x%2 == 0 {
				for y := 0; y < mylib.GetIntItem(values[x]); y++ {
					block = fmt.Sprintf("%s%d", block, index)
				}
				index++
			} else {
				for y := 0; y < mylib.GetIntItem(values[x]); y++ {
					block = fmt.Sprintf("%s%s", block, ".")
				}
			}
		}
	}

	blockS := strings.Split(block, "")
	pointL := 0
	pointR := len(block) - 1

	for pointL < pointR {
		// if blockS[pointR] == "." {
		// 	pointR--
		// } else if blockS[pointL] != "." {
		// 	pointL++
		// } else if blockS[pointL] == "." {
		// 	blockS[pointL] = blockS[pointR]
		// 	blockS[pointR] = "."
		// 	pointL++
		// 	pointR--
		// }
		if blockS[pointL] == "." {
			// fmt.Println(blockS)
			for blockS[pointR] == "." {
				// fmt.Println("Dot PointR : ", blockS[pointR], " : Place : ", pointR)
				pointR--
			}
			// fmt.Println("PointR : ", blockS[pointR], " : Place : ", pointR, " : L : ", pointL)
			blockS[pointL] = blockS[pointR]
			blockS[pointR] = "."
			pointR--
		}
		pointL++
	}

	// fmt.Println(blockS)
	pos := 0
	sum := 0
	for blockS[pos] != "." {
		sum += mylib.GetIntItem(blockS[pos]) * pos
		pos++
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
