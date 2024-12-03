package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x // control via return statement is a bad pattern. is there a better way to do this in go?
}

func dayOne() {

	file, err := os.Open("input-data/day1_input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	var vals1, vals2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("WARN: invalid line: ", line)
			continue
		}
		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("WARN: invalid data in line: ", line)
		}

		//fmt.Println(num1, num2)
		vals1 = append(vals1, num1)
		vals2 = append(vals2, num2)
	}

	sort.Ints(vals1)
	sort.Ints(vals2)

	var difs []int
	for i := range vals1 {
		difs = append(difs, vals2[i]-vals1[i])
		//fmt.Println(dif)
	}

	sum := 0
	for _, x := range difs {
		sum += abs(x)
	}
	fmt.Println("INFO: the differences sum to ", sum)

}
