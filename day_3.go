package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func extractPeskyDigits(s string) (int, int) {

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)
	digits := re.FindAllStringSubmatch(s, -1)[0]
	a, err1 := strconv.Atoi(digits[1])
	b, err2 := strconv.Atoi(digits[2])
	if err1 != nil || err2 != nil {
		log.Fatal("especially pesky digits detected in", digits)
	}
	return a, b
}
func findMulSubstrings(s string) []string {

	pattern := `mul\(\d{1,3},\d{1,3}\)`
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s, -1)
}

func dayThree() {

	lines, err := fileLineScanner("input-data/day3_input.txt")
	if err != nil {
		log.Fatal("Unable to read input file")
	}
	inputString := strings.Join(lines, "\n")
	mulStrings := findMulSubstrings(inputString)
	if len(mulStrings) == 0 {
		fmt.Println("WARN: no valid mul(XXX,YYY) substrings found")
	}
	var result []int
	for _, v := range mulStrings {
		a, b := extractPeskyDigits(v)
		result = append(result, a*b)
	}

	fmt.Println("result: ", sumIntArray(result))
}

func dayThreePartTwo() {

}
