package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func readDriveMap(lines string) []int {

	var result []int
	input := strings.TrimSpace(lines)
	block := true
	blockIndex := 0
	for _, c := range input {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal("ERROR: failed strcon.Atoi on: ", string(c))
		}
		if block {
			for i := 0; i < num; i++ {
				result = append(result, blockIndex)
			}
			blockIndex += 1
			block = false
		} else { // block == false
			for i := 0; i < num; i++ {
				result = append(result, -1)
			}
			block = true
		}
	}
	return result
}

func compressDrive(drive []int) { // TODO: test

	for i := 0; i < len(drive); i++ {
		if drive[i] == -1 {
			for j := len(drive) - 1; j > i; j-- {
				if drive[j] != -1 {
					drive[i], drive[j] = drive[j], drive[i] // swap
					break
				}
			}
		}
	}
}

func driveChecksum(drive []int) int {

	sum := 0
	for i, v := range drive {
		if v == -1 {
			break
		}
		sum += i * v
	}
	return sum
}

func printDriveMap(driveMap []int) {
	for _, v := range driveMap {
		if v == -1 {
			print(".")
		} else {
			print(v)
		}
	}
	print("\n")
}

func dayNine() {

	//lines := fileLineScanner("input-data-test/day9_input_test.txt")
	lines := fileLineScanner("input-data/day9_input.txt")

	d := readDriveMap(lines[0])
	compressDrive(d)
	fmt.Println("drive checksum: ", driveChecksum(d))
}
