package main

import (
	"bufio"
	"log"
	"os"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// count occurrences of all values in an integer array -> map
func countOccurrences(arr []int) map[int]int {

	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}
	return countMap
}

func fileLineScanner(filename string) []string {

	var values []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		values = append(values, line)
	}

	return values
}

func sumIntArray(arr []int) int {

	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}

func contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
