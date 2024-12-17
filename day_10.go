package main

import (
	"fmt"
	"log"
	"strconv"
)

type TopoMap struct {
	data  [][]int
	xSize int
	ySize int
}

// validDirections determines what directions the trail can take from the last Pos in path
// checks for edges and intersection with the existing path
func (t *TopoMap) validDirections(path []Pos) []Pos {

	directions := []Pos{
		Pos{1, 0},
		Pos{-1, 0},
		Pos{0, 1},
		Pos{0, -1},
		//Pos{1, 1},
		//Pos{-1, 1},
		//Pos{1, -1},
		//Pos{-1, -1},
	}
	if len(path) < 1 {
		log.Fatal("ERROR: validDirections: len(path) < 1: ", path)
	}
	current := path[len(path)-1]
	var result []Pos
	for _, dir := range directions {
		test := current.Add(dir)
		// if test Pos is off grid or is in the current path
		if test.y < 0 || test.y >= t.ySize || test.x < 0 || test.x >= t.xSize || contains(path, test) {
			continue
		} else if t.data[current.y][current.x] == t.data[test.y][test.x]-1 { // next Pos is on grid and is not on current path
			result = append(result, test)
		}
	}
	return result
}

func (t *TopoMap) countTrails(path []Pos) int {

	sum := 0
	validDirections := t.validDirections(path)
	currentPos := path[len(path)-1]
	currentHeight := t.data[currentPos.y][currentPos.x]
	//fmt.Println("INFO: countTrails ", currentPos, " ", currentHeight, " ", validDirections)
	if currentHeight == 9 {
		return 1
	}
	for _, nextPos := range validDirections {
		sum += t.countTrails(append(path, nextPos))
	}

	return sum
}

func (t *TopoMap) scoreTrails(visitedNines *[]Pos, path []Pos) int {

	sum := 0
	validDirections := t.validDirections(path)
	currentPos := path[len(path)-1]
	currentHeight := t.data[currentPos.y][currentPos.x]
	//fmt.Println("INFO: scoreTrails ", currentPos, " ", currentHeight, " ", validDirections)
	if currentHeight == 9 && !contains(*visitedNines, currentPos) {
		//fmt.Println("INFO: visitedNines: ", visitedNines)
		*visitedNines = append(*visitedNines, currentPos)
		return 1
	}
	for _, nextPos := range validDirections {
		sum += t.scoreTrails(visitedNines, append(path, nextPos))
	}

	return sum
}

func (t *TopoMap) TrailSearch() (int, int) {

	sum1 := 0
	sum2 := 0
	for y := range t.data {
		for x := range t.data[y] {
			if t.data[y][x] == 0 {
				//fmt.Printf("INFO: starting scoreTrails at [%d,%d] \n", x, y)
				visitedNines := []Pos{}
				sum1 += t.scoreTrails(&visitedNines, []Pos{Pos{x, y}})
				sum2 += t.countTrails([]Pos{Pos{x, y}})
				//fmt.Printf("INFO: finished scoreTrails at [%d,%d]: %d %d \n", x, y, sum1, sum2)
			}
		}
	}
	return sum1, sum2
}

func NewTopoMap(lines []string) TopoMap {

	result := TopoMap{}
	result.xSize = len(lines[0])
	result.ySize = len(lines)
	for _, line := range lines {
		var arr []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal("ERROR: invalid character in Atoi conversion: ", c)
			}
			arr = append(arr, num)
		}
		result.data = append(result.data, arr)
	}
	return result
}

func dayTen() {

	lines := fileLineScanner("input-data/day10_input.txt")
	//lines := fileLineScanner("input-data-test/day10_input_test.txt")

	topoMap := NewTopoMap(lines)
	//for _, r := range topoMap.data {
	//	fmt.Println(r)
	//}
	nTrailsPart1, nTrailsPart2 := topoMap.TrailSearch()
	fmt.Println("number of trails: ", nTrailsPart1)
	fmt.Println("number of paths: ", nTrailsPart2)
}
