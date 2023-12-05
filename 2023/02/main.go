package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}
GameLoop:
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		gameParts := strings.Split(parts[0], " ")
		id, _ := strconv.Atoi(gameParts[1])
		setParts := strings.Split(parts[1], "; ")

		for _, setPart := range setParts {
			pairParts := strings.Split(setPart, ", ")
			for _, pairPart := range pairParts {
				cubeParts := strings.Split(pairPart, " ")
				quantity, _ := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]
				if quantity > limits[color] {
					continue GameLoop
				}
			}
		}
		sum += id
	}
	fmt.Println(sum)
}

func Part2(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")
		setParts := strings.Split(parts[1], "; ")

		max := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, setPart := range setParts {
			pairParts := strings.Split(setPart, ", ")
			for _, pairPart := range pairParts {
				cubeParts := strings.Split(pairPart, " ")
				quantity, _ := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]
				if quantity > max[color] {
					max[color] = quantity
				}
			}
		}
		product := 1
		for _, v := range max {
			product *= v
		}
		sum += product
	}
	fmt.Println(sum)
}

func main() {
	part, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	filename := os.Args[2]
	switch part {
	case 1:
		Part1(filename)
	case 2:
		Part2(filename)
	default:
		fmt.Fprintf(os.Stderr, "Error: %v\n", "Invalid part")
		os.Exit(1)
	}
}
