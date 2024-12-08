package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	CX int
	CY int
}

func searchGrid2(grid [][]rune, start []int) int {
	directions := [][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	x, y := start[0], start[1]
	dir := 4
	d := directions[0]
	visited := map[[3]int]bool{}

	for {
		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
			break
		}

		state := [3]int{x, y, dir % 4}
		if visited[state] {
			return 1
		}
		visited[state] = true

		if grid[y][x] == '#' || grid[y][x] == 'O' {
			x -= d[0]
			y -= d[1]
			dir++
			mod := dir % 4
			d = directions[mod]
		}

		// fmt.Println("-----")
		// fmt.Println(x, y)
		// fmt.Println(visited[state])
		x += d[0]
		y += d[1]
	}

	return 0
}

func searchGrid(grid [][]rune, start []int) (int, []Coord) {
	directions := map[int][2]int{
		0: {0, -1},
		1: {1, 0},
		2: {0, 1},
		3: {-1, 0},
	}

	x, y := start[0], start[1]
	sum := 0
	dirCount := 4
	d := directions[0]

	path := []Coord{}
	for {
		if y < len(grid) && x < len(grid[0]) && x >= 0 && y >= 0 {
			if grid[y][x] == '#' || grid[y][x] == 'O' {
				x -= d[0]
				y -= d[1]
				dirCount++
				mod := dirCount % 4
				d = directions[mod]
			}
			if !(grid[y][x] == 'X') {
				grid[y][x] = 'X'
				path = append(path, Coord{CX: x, CY: y})
				sum++
			}

			x += d[0]
			y += d[1]
		} else {
			break
		}
	}

	return sum, path
}

func p2(grid [][]rune, start []int) {
	_, path := searchGrid(grid, start)
	sum := 0
	for coord := range path {
		if path[coord].CY == start[1] && path[coord].CX == start[0] {
			continue
		}
		fmt.Println(coord)
		grid[path[coord].CY][path[coord].CX] = 'O'
		sum += searchGrid2(grid, start)
		grid[path[coord].CY][path[coord].CX] = '.'
	}

	fmt.Println("P2:", sum)
}

func p1(grid [][]rune, start []int) {
	sum, _ := searchGrid(grid, start)
	fmt.Println("P1:", sum)
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
	var start []int
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		grid = append(grid, []rune(line))
		if strings.Contains(line, "^") {
			start = append(start, strings.Index(line, "^"))
			start = append(start, len(grid)-1)
		}
	}

	// p1(grid, start)
	p2(grid, start)
}
