package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkWord(x, y, dx, dy int, target string, grid [][]rune) bool {
	for i := 0; i < len(target); i++ {
		nx, ny := x+i*dx, y+i*dy
		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != rune(target[i]) {
			return false
		}
	}
	return true
}

func p1(grid [][]rune) {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	target := "XMAS"
	sum := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for _, dir := range directions {
				if checkWord(i, j, dir[0], dir[1], target, grid) {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}

func checkDiags(grid [][]rune, i, j int) int {
	directions := [][2]int{
		{-1, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
	}

	s := ""
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		nx1, ny1 := i+dx, j+dy

		if nx1 < 0 || ny1 < 0 || ny1 >= len(grid[0]) || nx1 >= len(grid) {
			return 0
		}

		if grid[nx1][ny1] == 'M' {
			s += "M"
		} else if grid[nx1][ny1] == 'S' {
			s += "S"
		}
	}

	if (s == "SMSM") || (s == "MSMS") || (s == "MSSM") || (s == "SMMS") {
		return 1
	}
	return 0
}

func p2(grid [][]rune) {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'A' {
				sum += checkDiags(grid, i, j)
			}
		}
	}

	fmt.Println("2.", sum)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		grid = append(grid, []rune(line))
	}

	p1(grid)
	p2(grid)
}
