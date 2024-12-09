package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Operator string

const (
	ADD Operator = "+"
	MUL Operator = "*"
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

	if len(ops) != len(rv)-1 {
		fmt.Println(rv)
		fmt.Println(ops)
		log.Fatal("ERROR: applyOps: length mismatch in rv/ops")
	}
	if len(rv) < 2 {
		fmt.Println(rv)
		fmt.Println(ops)
		log.Fatal("ERROR: applyOps: rv must contain at least two values")
	}
	sum := rv[0]
	for i := 0; i < len(ops); i++ {
		if ops[i] == ADD {
			sum += rv[i+1]
		} else { // ops[i] == MUL
			sum *= rv[i+1]
		}
	}
	return sum
}

func generatePermutations(n int, current []Operator, result *[][]Operator) { // TODO: grok this harder
	if len(current) == n {
		// Make a copy of the current slice and append it to the result
		temp := make([]Operator, n)
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	generatePermutations(n, append(current, ADD), result)
	generatePermutations(n, append(current, MUL), result)
}

// validEquation determines if the given equation is valid with any combination of operators
func validEquation(equation Equation) bool {

	var ops []Operator
	var celing []Operator
	for i := 0; i < len(equation.rv)-1; i++ {
		ops = append(ops, ADD)
		celing = append(celing, MUL)
	}
	//floorSum := applyOps(equation.rv, ops)
	//celingSum := applyOps(equation.rv, celing)

	// TODO: show peyton the cursed code below
	//if applyOps(equation.rv, ops) > equation.lv || applyOps(equation.rv, celing) < equation.lv {
	//	return false
	//}

	//if floorSum > equation.lv || celingSum < equation.lv {  // TODO: grok why this breaks the solution
	//	return false
	//}

	var operatorPermutations [][]Operator
	generatePermutations(len(ops), []Operator{}, &operatorPermutations)
	if len(operatorPermutations) != intPow(2, len(ops)) {
		fmt.Println(equation)
		fmt.Println(ops)
		log.Fatal("ERROR: something went wrong in generatePermutations")
	}
	// TODO: verify operatorPermutations contains only unique []Operator
	for _, p := range operatorPermutations {
		if applyOps(equation.rv, p) == equation.lv {
			return true
		}
	}
	return false

	//permute ops
	//for i := 0; i < len(ops); i++ {
	//	ops[i] = MUL
	//	if applyOps(equation.rv, ops) == equation.lv {
	//		return true
	//	}
	//	for j := i + 1; j < len(ops[j:]); j++ {
	//
	//	}
	//}
}

func daySeven() {

	lines := fileLineScanner("input-data/day7_input.txt")
	//lines := fileLineScanner("input-data-test/day7_input_test.txt")
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
