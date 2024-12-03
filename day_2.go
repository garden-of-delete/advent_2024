package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwo() {

	lines, err := fileLineScanner("input-data/day2_input.txt")
	if err != nil {
		return
	}

	var vals [][]int
	for _, line := range lines {
		var rowVals []int
		row := strings.Fields(line)
		for _, x := range row {
			num, err := strconv.Atoi(x)
			if err != nil {
				fmt.Println("WARN: invalid value in line:", line)
				continue
			}
			rowVals = append(rowVals, num)
		}
		vals = append(vals, rowVals)
	}

	safeReports := 0
	for _, row := range vals {
		// determine increasing or decreasing
		increasing := false
		safeRow := true
		if len(row) < 2 {
			fmt.Println("WARN: row with single value", row)
			continue // next row
		} else if row[0]-row[1] < 0 { // doesn't handle the case where row[0]-row[1] == 0
			increasing = true
		}
		for j := 0; j < len(row)-1; j++ { // for each value in the row
			// if increasing == true and (not increasing || increasing by > 3)
			if increasing && (row[j]-row[j+1] >= 0 || row[j]-row[j+1] < -3) {
				//fmt.Println("INFO: unsafe increasing row", row)
				safeRow = false
				break // next row
				// if increasing == false and (not decreasing || decreasing by > 3)
			} else if !increasing && (row[j]-row[j+1] <= 0 || row[j]-row[j+1] > 3) {
				//fmt.Println("INFO: unsafe decreasing row", row)
				safeRow = false
				break
			}
		}
		if safeRow {
			safeReports += 1
		}
	}

	fmt.Printf("INFO: %d safe reports\n", safeReports)
}
