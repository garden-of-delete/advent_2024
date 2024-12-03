package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func dayOne() {

	lines, err := fileLineScanner("input-data/day1_input.txt")
	if err != nil {
		return
	}

	var vals1, vals2 []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("WARN: invalid line: ", line)
			continue
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("WARN: invalid data in line: ", line)
			continue
		}
		vals1 = append(vals1, num1)
		vals2 = append(vals2, num2)

	}

	sort.Ints(vals1)
	sort.Ints(vals2)

	var difs []int
	for i := range vals1 {
		difs = append(difs, abs(vals2[i]-vals1[i]))
		//fmt.Println(dif)
	}

	fmt.Println("INFO: the differences sum to ", sumIntArray(difs))

	var similarity []int
	countMap := countOccurrences(vals2)
	for i := range vals1 {
		if v, ok := countMap[vals1[i]]; ok {
			similarity = append(similarity, vals1[i]*v)
		} else {
			continue
		}
	}
	fmt.Println("INFO: similarity score ", sumIntArray(similarity))
}
