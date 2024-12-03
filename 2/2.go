package main

import (
	"1/mylib"
	"fmt"
	"log"

	// "sort"
	// "strconv"
	"strings"
)

func main() {
	scanner, file := mylib.GetFileScanner()

	safeReports := 0
	safeReportsTol1 := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		safeExit := isSafe(items, 0)
		safeExitTol1 := isSafe(items, 1)
		if safeExit {
			fmt.Println("Safe Exit : ", line)
			safeReports++
		}
		if safeExitTol1 {
			fmt.Println("Safe Tol1 Exit : ", line)
			safeReportsTol1++
		}
	}

	fmt.Println("Safe Report: ", safeReports)
	fmt.Println("Safe Report Tolerance 1: ", safeReportsTol1)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

func isSafe(items []string, tolerance int) bool {
	isIncreasing := false
	isDecreasing := false
	safeExit := false
	tolerateCount := 0
	for x := 1; x < len(items); x++ {
		diff := mylib.GetIntDiff(mylib.GetIntItem(items[x]), mylib.GetIntItem(items[x-1]))
		if diff > 3 {
			// fmt.Println("Three Exit : ", line, diff, lastDiff)
			safeExit = false
			if tolerateCount == tolerance {
				break
			}
			tolerateCount++
		}
		if diff == 0 {
			// fmt.Println("Same Exit : ", line, diff, lastDiff)
			safeExit = false
			if tolerateCount == tolerance {
				break
			}
			tolerateCount++
		}
		if mylib.GetIntItem(items[x]) > mylib.GetIntItem(items[x-1]) {
			isIncreasing = true
			if isDecreasing {
				// fmt.Println("Dec Exit : ", line)
				safeExit = false
				if tolerateCount == tolerance {
					break
				}
				tolerateCount++
			}
		}
		if mylib.GetIntItem(items[x]) < mylib.GetIntItem(items[x-1]) {
			isDecreasing = true
			if isIncreasing {
				// fmt.Println("Inc Exit : ", line)
				safeExit = false
				if tolerateCount == tolerance {
					break
				}
				tolerateCount++
			}
		}
		safeExit = true
	}
	return safeExit
}
