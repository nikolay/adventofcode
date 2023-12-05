package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsDigit(grid []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(grid) {
		return false
	}
	if row := grid[y]; x < len(row) {
		c := row[x]
		return c >= '0' && c <= '9'
	}
	return false
}

func Part1(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if c := grid[y][x]; c != '.' && (c < '0' || c > '9') {
				for cy := y - 1; cy < y+2; cy++ {
					for cx := x - 1; cx < x+2; cx++ {
						if (cx != x || cy != y) && IsDigit(grid, cx, cy) {
							var l, r int
							for l = cx; IsDigit(grid, l-1, cy); l-- {
							}
							for r = cx; IsDigit(grid, r, cy); r++ {
							}
							n, _ := strconv.Atoi((grid[cy])[l:r])
							sum += n
							grid[cy] = grid[cy][:l] + strings.Repeat(".", r-l) + grid[cy][r:]
						}
					}
				}
			}
		}
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

	grid := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if c := grid[y][x]; c == '*' {
				count := 0
				product := 1
				for cy := y - 1; cy < y+2; cy++ {
					for cx := x - 1; cx < x+2; cx++ {
						if (cx != x || cy != y) && IsDigit(grid, cx, cy) {
							var l, r int
							for l = cx; IsDigit(grid, l-1, cy); l-- {
							}
							for r = cx; IsDigit(grid, r, cy); r++ {
							}
							n, _ := strconv.Atoi((grid[cy])[l:r])
							product = product * n
							count++
							grid[cy] = grid[cy][:l] + strings.Repeat(".", r-l) + grid[cy][r:]
						}
					}
				}
				if count == 2 {
					sum += product
				}
			}
		}
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
