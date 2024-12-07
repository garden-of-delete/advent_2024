package main

import (
	"fmt"
	"strconv"
	"strings"
)

func swapLeft[T any](arr []T, i int) {

	if i == 0 { // if not the first index
		return
	}
	arr[i], arr[i-1] = arr[i-1], arr[i]
}

func getInvalidIndex(arr []int, orderMap map[int]*Set[int]) int {

	seen := NewSet[int]()
	for i, v := range arr { // for each value in the sequence
		if i == 0 { // if the first value in the sequence
			seen.Add(v)
			continue
		} else if _, exists := orderMap[v]; exists { // if any other values must appear before this one
			if seen.Intersects(orderMap[v]) {
				return i
			}
			seen.Add(v)
		} else {
			seen.Add(v)
		}
	}
	return -1
}

func orderPages(arr []int, orderMap map[int]*Set[int]) []int {

	var ordered []int
	for _, v := range arr { // deep copy input array
		ordered = append(ordered, v)
	}
	for i := getInvalidIndex(ordered, orderMap); i != -1; i = getInvalidIndex(ordered, orderMap) {
		swapLeft(ordered, i)
	}
	return ordered
}

func dayFive() {

	//lines := fileLineScanner("input-data-test/day5_input_test.txt")
	lines := fileLineScanner("input-data/day5_input.txt")
	var orderRules, pageSequence [][]int
	for _, v := range lines {
		if strings.Contains(v, "|") {
			vals := strings.Split(v, "|")
			num1, err1 := strconv.Atoi(vals[0])
			num2, err2 := strconv.Atoi(vals[1])
			if err1 != nil || err2 != nil {
				fmt.Println("WARN: invalid integers in row ", v)
				continue
			}
			orderRules = append(orderRules, []int{num1, num2}) // num1 MUST appear before num2
		} else {
			var nums []int
			vals := strings.Split(v, ",")
			for _, w := range vals {
				num, err := strconv.Atoi(w)
				if err != nil {
					fmt.Println("WARN: invalid integers in row ", v)
					continue
				}
				nums = append(nums, num)
			}
			pageSequence = append(pageSequence, nums)
		}
	}

	// populate order rules as a map from int -> Set[int]
	orderMap := make(map[int]*Set[int]) // k must be before {v,...}
	for i := range orderRules {
		_, exists := orderMap[orderRules[i][0]]
		if !exists {
			orderMap[orderRules[i][0]] = NewSet[int]()
		}
		orderMap[orderRules[i][0]].Add(orderRules[i][1])
	}

	// orderMap -> key MUST appear before all values in the value Set
	nInvalidOrderings := 0
	var fixedOrderings [][]int
	var validOrderingMiddleValues []int
	for _, arr := range pageSequence { // for each sequence of pages
		isValidOrdering := true
		seen := NewSet[int]()
		for i, v := range arr { // for each value in the sequence
			if i == 0 { // if the first value in the sequence
				seen.Add(v)
				continue
			} else if _, exists := orderMap[v]; exists { // if any other values must appear before this one
				if seen.Intersects(orderMap[v]) {
					//fmt.Println("INFO: Invalid ordering: ", arr)
					isValidOrdering = false
					nInvalidOrderings++
					fixedOrderings = append(fixedOrderings, orderPages(arr, orderMap))
					break
				}
			}
			seen.Add(v)
		}
		if isValidOrdering {
			validOrderingMiddleValues = append(validOrderingMiddleValues, arr[(len(arr)-1)/2])
		}
	}

	//fmt.Println("invalid orderings: ", nInvalidOrderings)
	fmt.Println("valid ordering middle value sum: ", sumIntArray(validOrderingMiddleValues))
	//fmt.Println("fixed orderings: ", fixedOrderings)

	var fixedOrderingMiddleValues []int
	for _, arr := range fixedOrderings {
		fixedOrderingMiddleValues = append(fixedOrderingMiddleValues, arr[(len(arr)-1)/2])
	}
	fmt.Println("fixed ordering middle value sum: ", sumIntArray(fixedOrderingMiddleValues))
}
