package main

import (
	"1/mylib"
	"fmt"
	"log"
	"strings"
)

type Cord struct {
	x, y interface{}
}

func main() {
	scanner, file := mylib.GetFileScanner()

	radarsCord := make(map[string][]Cord)

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, "")
		lineNumber := 0
		for x := 0; x < len(points); x++ {
			if points[x] != "." {
				radarsCord[points[x]] = append(radarsCord[points[x]], Cord{x, lineNumber})
			}
		}
		lineNumber++
	}
	diffCords := make(map[string][]Cord)
	for k, v := range radarsCord {
		for x := 0; x < len(radarsCord[k]); x++ {
			for y := x; y < len(radarsCord[k]); y++ {
				diffCords[k] = append(diffCords[k], Cord{v[x].x.(int) - v[y].x.(int), v[x].y.(int) - v[y].y.(int)})
			}
		}
	}
	fmt.Println(radarsCord)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}
