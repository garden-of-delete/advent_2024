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

func (a Pos) Scale(n int) Pos {

	return Pos{a.x * n, a.y * n}
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

func findAntiNodes(cityMap *CityMap, pair PosPair) []Pos {

	var result []Pos
	vector := pair.b.Sub(pair.a) // a -> b
	for i := 0; ; i++ {
		p := pair.a.Add(vector.Scale(i))
		if isOnMap(cityMap, p) {
			result = append(result, p)
		} else {
			break
		}
	}
	for i := -1; ; i-- {
		p := pair.a.Add(vector.Scale(i))
		if isOnMap(cityMap, p) {
			result = append(result, p)
		} else {
			break
		}
	}
	return result
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
			t := findAntiNodes(&cityMap, pair) // on-map anti-nodes only
			for _, a := range t {
				result[a] = frequency
			}
		}
	}
	return result
}

func dayEight() {

	lines := fileLineScanner("input-data/day8_input.txt")

	cityMap := readCityMap(lines)
	//fmt.Println("city map: ")
	//printCityMap(cityMap)

	//fmt.Println("=== === === ===")

	antiNodes := scanFrequencies(cityMap)
	nAntiNodes := 0
	for y := 0; y < cityMap.ySize; y++ {
		for x := 0; x < cityMap.xSize; x++ {
			if _, exists := antiNodes[Pos{x, y}]; exists {
				//print("#")
				nAntiNodes++
			} else {
				//print(".")
			}
		}
		//print("\n")
	}
	fmt.Println("found antinodes: ", nAntiNodes)

}
