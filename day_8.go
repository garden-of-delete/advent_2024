package main

// use Pos from Day 6
//type Pos struct {
//	x int
//	y int
//}

type CityMap struct {
	antennas map[string][]Pos
	positions map[Pos]string
	xSize    int
	ySize    int
}

func readCityMap(lines []string) CityMap {

	xSize := len(lines[0])
	ySize := len(lines)
	var antennas map[string][]Pos
	var positions map[Pos]string

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
				antennas[s] = append(antennas[s], Pos{x,y})
			}
		}
	}
	return CityMap{antennas, positions, xSize, ySize}
}

func printCityMap(cityMap CityMap) {

	for y := 0; y < cityMap.ySize; y++ {
		for x:= 0; x < cityMap.xSize; x++ {
			for a := range cityMap.antennas
		}
	}
}

func dayEight() {

	//lines := fileLineScanner("input-data/day8_input.txt")
	lines := fileLineScanner("input-data-test/day8_input_test.txt")

	cityMap := readCityMap(lines)

}
