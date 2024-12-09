package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operator string

const (
	ADD    Operator = "+"
	MUL    Operator = "*"
	CONCAT Operator = "||"
)

type Equation struct {
	lv int
	rv []int
}

func ParseEquations(inputs []string) ([]Equation, error) {

	var equations []Equation
	for _, input := range inputs {

		parts := strings.Split(input, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid input format: %s", input)
		}

		lv, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid left value in '%s': %w", input, err)
		}

		rightParts := strings.Split(parts[1], ",")
		var rv []int
		for _, part := range rightParts {
			vals := strings.Split(part, " ")
			for _, v := range vals {
				if v != "" {
					num, err := strconv.Atoi(v)
					if err != nil {
						fmt.Printf("ERROR: unable to parse val %s in right part %s\n", v, part)
						return nil, fmt.Errorf(
							"ERROR: unable to parse val %s in right part %s: %w", v, part, err)
					} else {
						rv = append(rv, num)
					}
				}
			}
		}
		equations = append(equations, Equation{lv: lv, rv: rv})
	}
	return equations, nil
}

// applyOps applies the operations to rv
func applyOps(rv []int, ops []Operator) int {

	sum := rv[0]
	for i := 0; i < len(ops); i++ {
		if ops[i] == ADD {
			sum += rv[i+1]
		} else if ops[i] == MUL {
			sum *= rv[i+1]
		} else { // ops[i] == CONCAT
			str1 := strconv.Itoa(sum)
			str2 := strconv.Itoa(rv[i+1])
			concatenated := str1 + str2
			r, err := strconv.Atoi(concatenated)
			if err != nil {
				fmt.Println(rv)
				fmt.Println(ops)
				fmt.Println(concatenated)
				log.Fatal("ERROR: applyOps: cant convert concat string to int", err)
			} else {
				sum = r
			}
		}
	}
	return sum
}

func generatePermutations(n int, current []Operator, result *[][]Operator) {

	if len(current) == n {
		// Make a copy of the current slice and append it to the result
		temp := make([]Operator, n)
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	generatePermutations(n, append(current, ADD), result)
	generatePermutations(n, append(current, MUL), result)
	generatePermutations(n, append(current, CONCAT), result)
}

// validEquation determines if the given equation is valid with any combination of operators
func validEquation(equation Equation) bool {

	var ops []Operator
	for i := 0; i < len(equation.rv)-1; i++ {
		ops = append(ops, ADD)
	}

	var operatorPermutations [][]Operator
	generatePermutations(len(ops), []Operator{}, &operatorPermutations)
	if len(operatorPermutations) != intPow(3, len(ops)) {
		fmt.Println(equation)
		fmt.Println(ops)
		log.Fatal("ERROR: something went wrong in generatePermutations")
	}
	for _, p := range operatorPermutations {
		if applyOps(equation.rv, p) == equation.lv {
			return true
		}
	}
	return false
}

func daySeven() {

	lines := fileLineScanner("input-data/day7_input.txt")
	equations, err := ParseEquations(lines)
	if err != nil {
		log.Fatal("invalid input data. ", err)
	}

	var validEquations []Equation
	for _, e := range equations {
		if validEquation(e) {
			validEquations = append(validEquations, e)
		}
	}

	validEquationsSum := 0
	for _, e := range validEquations {
		validEquationsSum += e.lv
	}
	fmt.Println("valid equations sum: ", validEquationsSum)
}
