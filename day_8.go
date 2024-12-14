package main

import (
	"fmt"
)

// use Pos from Day 6
//type Pos struct {
//	x int
//	y int
//}

type CityMap struct {
	antennas map[string][]Pos
	xSize    int
	ySize    int
}

type PosPair struct {
	a Pos
	b Pos
}

// returns a + b
func (a Pos) Add(b Pos) Pos {

	return Pos{a.x + b.x, a.y + b.y}
}

// returns a - b
func (a Pos) Sub(b Pos) Pos {

	return Pos{a.x - b.x, a.y - b.y}
}

func readCityMap(lines []string) CityMap {

	xSize := len(lines[0])
	ySize := len(lines)
	antennas := map[string][]Pos{}

	for y := range lines {
		chars := []rune(lines[y])
		for x := range chars {
			if chars[x] == '.' {
				continue
			} else { // chars[x] is alphanumeric character
				s := string(chars[x])
				if _, exists := antennas[s]; !exists {
					antennas[s] = []Pos{}
				}
				antennas[s] = append(antennas[s], Pos{x, y})
			}
		}
	}
	return CityMap{antennas, xSize, ySize}
}

func printCityMap(cityMap CityMap) {

	for y := 0; y < cityMap.ySize; y++ {
		for x := 0; x < cityMap.xSize; x++ {
			found := false
			for a := range cityMap.antennas {
				for _, pos := range cityMap.antennas[a] {
					if pos.x == x && pos.y == y {
						print(a)
						found = true
						break
					}
				}
				if found == true {
					break
				}
			}
			if !found {
				print(".")
			}
		}
		print("\n")
	}
}

func findAntiNodes(pair PosPair) (Pos, Pos) {
	// pair.a=(1,3) pair.b=(5,9) -> (4,6)
	// antinode 1 = (1,3)+(-4,-6)=(-3,-3)
	// antinode 2 = (5,9)-(-4,-6)=(9,15)
	// (5,9) (1,3) -> (4,6)
	vector := pair.b.Sub(pair.a)
	antiNode1 := pair.a.Sub(vector)
	antiNode2 := pair.b.Add(vector)
	return antiNode1, antiNode2
}

func isOnMap(cityMap *CityMap, pos Pos) bool {

	if pos.x < 0 || pos.x >= cityMap.xSize || pos.y < 0 || pos.y >= cityMap.ySize {
		return false
	}
	return true
}

func scanFrequencies(cityMap CityMap) map[Pos]string {

	result := map[Pos]string{}
	for frequency := range cityMap.antennas {
		var pairs []PosPair // TODO: is a null map?
		// compute all pairwise combinations of antennas on this frequency
		for i := 0; i < len(cityMap.antennas[frequency])-1; i++ {
			for j := i + 1; j < len(cityMap.antennas[frequency]); j++ {
				pairs = append(pairs, PosPair{cityMap.antennas[frequency][i], cityMap.antennas[frequency][j]})
			}
		}
		// find anti-nodes and store the ones that are on map
		for _, pair := range pairs { // for each pair at this frequency
			a, b := findAntiNodes(pair)
			if isOnMap(&cityMap, a) {
				result[a] = frequency
			}
			if isOnMap(&cityMap, b) {
				result[b] = frequency
			}
		}
	}
	return result
}

func dayEight() {

	lines := fileLineScanner("input-data/day8_input.txt")
	//lines := fileLineScanner("input-data-test/day8_input_test.txt")

	cityMap := readCityMap(lines)
	fmt.Println("city map: ")
	printCityMap(cityMap)

	fmt.Println("=== === === ===")
	//a := Pos{1, 3}
	//b := Pos{5, 9}
	//test1 := PosPair{a, b}
	//test2 := PosPair{b, a}
	//fmt.Println(findAntiNodes(test1))
	//fmt.Println(findAntiNodes(test2))

	antiNodes := scanFrequencies(cityMap)
	nAntiNodes := 0
	for y := 0; y < cityMap.ySize; y++ {
		for x := 0; x < cityMap.xSize; x++ {
			if _, exists := antiNodes[Pos{x, y}]; exists {
				print("#")
				nAntiNodes++
			} else {
				print(".")
			}
		}
		print("\n")
	}
	fmt.Println("found antinodes: ", nAntiNodes)

}
