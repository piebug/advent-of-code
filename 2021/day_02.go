package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day02 struct {
	aimOn bool
}

func (d Day02) Solve(scanner *bufio.Scanner) {
	var x, y, aim int

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")

		direction := command[0]
		units, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case "forward":
			x += units
			if d.aimOn {
				y += units * aim
			}
		case "up":
			if d.aimOn {
				aim -= units
			} else {
				y -= units
			}
		case "down":
			if d.aimOn {
				aim += units
			} else {
				y += units
			}
		}
	}

	fmt.Printf("The final position is (%d, %d) and the product is: %d\n", x, y, x*y)
}
