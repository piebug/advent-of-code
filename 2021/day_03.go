package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day03 struct{}

func (d Day03) Solve(scanner *bufio.Scanner) {
	var bitCounter []int
	var total int

	for scanner.Scan() {
		bits := strings.Split(scanner.Text(), "")
		total++

		for i, bit := range bits {
			if bitCounter == nil || i >= len(bitCounter) {
				bitCounter = append(bitCounter, 0)
			}
			if bit == "1" {
				bitCounter[i]++
			}
		}
	}

	var gamma, epsilon string
	half := total / 2

	fmt.Printf("%d+, half: %d", bitCounter, half)

	for _, count := range bitCounter {
		if count > half {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	fmt.Printf("total: %d, gamma: %s, epsilon: %s", total, gamma, epsilon)

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Printf("gamma: %d, epsilon: %d", gammaInt, epsilonInt)

	fmt.Printf("The power consumption is: %d\n", gammaInt*epsilonInt)
}
