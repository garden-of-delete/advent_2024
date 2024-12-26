package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func readVals(lines []string) []int {

	var result []int
	t := strings.Split(lines[0], " ")
	for _, s := range t {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("unable to Atoi: ", num)
		}
		result = append(result, num)
	}
	return result
}

func evolve(stones []int) []int {

	//TODO: pre-allocate?
	//implement rules
	var result []int
	for i := range stones {
		if stones[i] == 0 {
			result = append(result, 1)
		} else if runes := []rune((strconv.Itoa(stones[i]))); len(runes)%2 == 0 {
			// do array operation
			left, err1 := strconv.Atoi(string(runes[:len(runes)/2]))
			right, err2 := strconv.Atoi(string(runes[len(runes)/2:]))
			if err1 != nil || err2 != nil {
				log.Fatal("ERROR: cant convert", left, right)
			}
			result = append(result, left)
			result = append(result, right)

		} else {
			result = append(result, stones[i]*2024)
		}
	}
	return result
}

func dayEleven() {

	lines := fileLineScanner("input-data/day11_input.txt")
	//lines := fileLineScanner("input-data-test/day11_input_test.txt")

	vals := readVals(lines)
	fmt.Println(vals)

	nEvolutions := 75
	for i := 0; i < nEvolutions; i++ {
		fmt.Println("INFO: evolve iteration ", i)
		vals = evolve(vals)
		fmt.Println("INFO: nStones: ", len(vals) /*, " ", vals*/)
	}
}
