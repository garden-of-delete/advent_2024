package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkReportIsValid(slice []int) bool {

	isIncreasing := true
	if slice[0]-slice[1] > 0 {
		isIncreasing = false
	}
	for i := 0; i < len(slice)-1; i++ {
		diff := slice[i] - slice[i+1]
		if diff == 0 {
			return false
		}
		if isIncreasing && (diff >= 0 || abs(diff) > 3) {
			return false
		} else if !isIncreasing && (diff <= 0 || abs(diff) > 3) {
			return false
		}
	}
	return true
}

func excise[T any](slice []T, i int) []T { // this is fine.... [fire intensifies]

	newSlice := make([]T, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	newSlice = append(newSlice, slice[i+1:]...)
	return newSlice
}

func reverseCopy(arr []int) []int {
	n := len(arr)
	reversed := make([]int, n) // Create a new slice of the same length
	for i := 0; i < n; i++ {
		reversed[i] = arr[n-i-1]
	}
	return reversed
}

type Report struct {
	originalIsValid bool // default value is false
	excisedIsValid  bool
}

func exciseScan(slice []int) Report { // it's pasta time
	reversed := reverseCopy(slice)
	//fmt.Println("for slice", slice)
	// check if report is valid in either direction
	if checkReportIsValid(slice) || checkReportIsValid(reversed) {
		//fmt.Println("original true, excised false")
		return Report{originalIsValid: true, excisedIsValid: false}
	} else {
		//fmt.Println("starting excisions...")
		for i := range slice {
			excised := excise(slice, i)
			//fmt.Println("excised slice", excised)
			reversed = reverseCopy(excised)
			if checkReportIsValid(excised) || checkReportIsValid(reversed) {
				//fmt.Printf("original false, excised true")
				return Report{originalIsValid: false, excisedIsValid: true}
			}
		}
		//fmt.Println("original false, excised false")
		return Report{originalIsValid: false, excisedIsValid: false}
	}
}

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
	safeReportsWithDampener := 0
	for _, row := range vals {
		// determine increasing or decreasing
		if len(row) < 2 {
			fmt.Println("WARN: row with single value", row)
			continue // next row
		}
		result := exciseScan(row)
		if result.originalIsValid {
			safeReports++
		} else if result.excisedIsValid {
			safeReportsWithDampener++
		}
	}

	fmt.Printf("INFO: %d safe reports\n", safeReports)
	fmt.Printf("INFO: %d safe reports w/ dampener\n", safeReports+safeReportsWithDampener)
}
