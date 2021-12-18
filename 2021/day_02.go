package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day02(aimOn bool) {
	file, err := os.Open("data/day_02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scanner by default splits at newline
	scanner := bufio.NewScanner(file)

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
			if aimOn {
				y += units * aim
			}
		case "up":
			if aimOn {
				aim -= units
			} else {
				y -= units
			}
		case "down":
			if aimOn {
				aim += units
			} else {
				y += units
			}
		}
	}

	fmt.Printf("The final position is (%d, %d) and the product is: %d\n", x, y, x*y)
}
