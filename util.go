package main

import (
	"bufio"
	"log"
	"os"
	"strings"
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

func fileLineScanner(filename string) ([]string, error) {

	var values []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return values, err
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
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		values = append(values, line)
	}

	return values, nil
}

func sumIntArray(arr []int) int {

	sum := 0
	for _, x := range arr {
		sum += x
	}
	return sum
}
