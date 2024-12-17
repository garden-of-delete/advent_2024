package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Drive struct {
	d           []int
	nFileBlocks int
}

func readDriveMap(lines string) Drive {

	var result Drive
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
				result.d = append(result.d, blockIndex)
			}
			blockIndex += 1
			block = false
		} else { // block == false
			for i := 0; i < num; i++ {
				result.d = append(result.d, -1)
			}
			block = true
		}
		result.nFileBlocks = blockIndex
	}
	return result
}

func (drive *Drive) findFileByIndex(fileIndex int) (startIndex, endIndex int) {

	for i := 0; i < len(drive.d); i++ {
		if drive.d[i] == fileIndex {
			startIndex = i
			for j := i; j < len(drive.d) && drive.d[j] == drive.d[i]; j++ {
				endIndex = j
			}
			break
		}
	}
	return
}

func (drive *Drive) findSpace(size int) (startIndex int) {

	for i := 0; i < len(drive.d); i++ {
		if drive.d[i] == -1 {
			for j := i; j < len(drive.d) && drive.d[j] == drive.d[i]; j++ {
				if j-i+1 >= size {
					return i
				}
			}
		}
	}
	return -1
}

func (drive *Drive) moveBlock(blockIndex, destination int) { // can corrupt drive if destination does not have enough space

	startIndex, endIndex := drive.findFileByIndex(blockIndex)
	for i := 0; i < endIndex-startIndex+1; i++ {
		drive.d[destination+i] = blockIndex
	}
	for i := startIndex; i <= endIndex; i++ {
		drive.d[i] = -1
	}
}

func compressDrive(drive []int) {

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

func defragCompressDrive(drive Drive) {

	for i := drive.nFileBlocks - 1; i >= 0; i-- {
		start, end := drive.findFileByIndex(i)
		if spaceStart := drive.findSpace(end - start + 1); spaceStart != -1 {
			if spaceStart < start {
				drive.moveBlock(i, spaceStart)
			}

		}
		//fmt.Println("defrag block ", i)
		//printDriveMap(drive.d)
	}
}

func driveChecksum(drive []int) int {

	sum := 0
	for i, v := range drive {
		if v != -1 {
			sum += i * v
		}
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

	lines := fileLineScanner("input-data/day9_input.txt")

	drive := readDriveMap(lines[0])
	//printDriveMap(drive.d)
	//compressDrive(drive.d)

	defragCompressDrive(drive)
	//printDriveMap(drive.d)
	fmt.Println("drive checksum: ", driveChecksum(drive.d))
}
